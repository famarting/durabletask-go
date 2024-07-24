package backend

import (
	context "context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/microsoft/durabletask-go/api"
	"github.com/microsoft/durabletask-go/internal/helpers"
	"github.com/microsoft/durabletask-go/internal/protos"
)

var emptyCompleteTaskResponse = &protos.CompleteTaskResponse{}

var errShuttingDown error = status.Error(codes.Canceled, "shutting down")

type ExecutionResults struct {
	Response *protos.OrchestratorResponse
}

type Executor interface {
	ExecuteOrchestrator(ctx context.Context, iid api.InstanceID, oldEvents []*protos.HistoryEvent, newEvents []*protos.HistoryEvent) (*ExecutionResults, error)
	ExecuteActivity(context.Context, api.InstanceID, *protos.HistoryEvent) (*protos.HistoryEvent, error)
	Shutdown(ctx context.Context) error
}

type grpcExecutor struct {
	protos.UnimplementedTaskHubSidecarServiceServer
	inflightWorkItems    *sync.Map
	backend              Backend
	logger               Logger
	onWorkItemConnection func(context.Context) error
	streamShutdownChan   <-chan any
}

type grpcExecutorOptions func(g *grpcExecutor)

// IsDurableTaskGrpcRequest returns true if the specified gRPC method name represents an operation
// that is compatible with the gRPC executor.
func IsDurableTaskGrpcRequest(fullMethodName string) bool {
	return strings.HasPrefix(fullMethodName, "/TaskHubSidecarService/")
}

// WithOnGetWorkItemsConnectionCallback allows the caller to get a notification when an external process
// connects over gRPC and invokes the GetWorkItems operation. This can be useful for doing things like
// lazily auto-starting the task hub worker only when necessary.
func WithOnGetWorkItemsConnectionCallback(callback func(context.Context) error) grpcExecutorOptions {
	return func(g *grpcExecutor) {
		g.onWorkItemConnection = callback
	}
}

func WithStreamShutdownChannel(c <-chan any) grpcExecutorOptions {
	return func(g *grpcExecutor) {
		g.streamShutdownChan = c
	}
}

// NewGrpcExecutor returns the Executor object and a method to invoke to register the gRPC server in the executor.
func NewGrpcExecutor(be Backend, logger Logger, opts ...grpcExecutorOptions) (executor Executor, registerServerFn func(grpcServer grpc.ServiceRegistrar)) {
	grpcExecutor := &grpcExecutor{
		backend:           be,
		logger:            logger,
		inflightWorkItems: &sync.Map{},
	}

	for _, opt := range opts {
		opt(grpcExecutor)
	}

	return grpcExecutor, func(grpcServer grpc.ServiceRegistrar) {
		protos.RegisterTaskHubSidecarServiceServer(grpcServer, grpcExecutor)
	}
}

// ExecuteOrchestrator implements Executor
func (executor *grpcExecutor) ExecuteOrchestrator(ctx context.Context, iid api.InstanceID, oldEvents []*protos.HistoryEvent, newEvents []*protos.HistoryEvent) (*ExecutionResults, error) {
	executor.logger.Debugf("executing orchestrator %v", string(iid))

	waitCtx, waitCancel := context.WithCancel(ctx)
	defer waitCancel()

	err := executor.backend.SetPendingOrchestrator(ctx, iid)
	if err != nil {
		return nil, err
	}

	key := getOrchestratorExecutionKey(string(iid))
	executor.inflightWorkItems.Store(key, waitCancel)
	defer executor.inflightWorkItems.Delete(key)

	workItem := &protos.WorkItem{
		Request: &protos.WorkItem_OrchestratorRequest{
			OrchestratorRequest: &protos.OrchestratorRequest{
				InstanceId:  string(iid),
				ExecutionId: nil,
				PastEvents:  oldEvents,
				NewEvents:   newEvents,
			},
		},
	}

	err = executor.backend.EnqueueWorkItem(ctx, workItem)
	if err != nil {
		if ctx.Err() != nil {
			executor.logger.Warnf("%s: context canceled before dispatching orchestrator work item", iid)
			return nil, ctx.Err()
		}
		executor.logger.Errorf("%s: error queuing work item %v", iid, err)
		return nil, err
	}

	// Wait for the connected worker to signal that it's done executing the work-item
	response, err := executor.backend.WaitForPendingOrchestrator(waitCtx, iid)
	if err != nil {
		if waitCtx.Err() != nil {
			return nil, errors.New("operation aborted")
		}
		if ctx.Err() != nil {
			executor.logger.Warnf("%s: context canceled before receiving orchestrator result", iid)
			return nil, ctx.Err()
		}
		return nil, err
	}
	executor.logger.Debugf("%s: orchestrator got result", iid)
	if response == nil {
		return nil, errors.New("operation aborted")
	}

	return &ExecutionResults{
		Response: response,
	}, nil
}

// ExecuteActivity implements Executor
func (executor *grpcExecutor) ExecuteActivity(ctx context.Context, iid api.InstanceID, e *protos.HistoryEvent) (*protos.HistoryEvent, error) {
	executor.logger.Debugf("executing activity %v %d", string(iid), e.EventId)

	waitCtx, waitCancel := context.WithCancel(ctx)
	defer waitCancel()

	err := executor.backend.SetPendingActivity(ctx, iid, e.EventId)
	if err != nil {
		return nil, err
	}

	key := getActivityExecutionKey(string(iid), e.EventId)
	executor.inflightWorkItems.Store(key, waitCancel)
	defer executor.inflightWorkItems.Delete(key)

	task := e.GetTaskScheduled()
	workItem := &protos.WorkItem{
		Request: &protos.WorkItem_ActivityRequest{
			ActivityRequest: &protos.ActivityRequest{
				Name:                  task.Name,
				Version:               task.Version,
				Input:                 task.Input,
				OrchestrationInstance: &protos.OrchestrationInstance{InstanceId: string(iid)},
				TaskId:                e.EventId,
			},
		},
	}

	err = executor.backend.EnqueueWorkItem(ctx, workItem)
	if err != nil {
		if ctx.Err() != nil {
			executor.logger.Warnf("%s/%s#%d: context canceled before dispatching activity work item", iid, task.Name, e.EventId)
			return nil, ctx.Err()
		}
		executor.logger.Errorf("%s/%s#%d: error queuing work item %v", iid, task.Name, e.EventId, err)
		return nil, err
	}

	// Wait for the connected worker to signal that it's done executing the work-item
	response, err := executor.backend.WaitForPendingActivity(waitCtx, iid, e.EventId)
	if err != nil {
		if waitCtx.Err() != nil {
			return nil, errors.New("operation aborted")
		}
		if ctx.Err() != nil {
			executor.logger.Warnf("%s/%s#%d: context canceled before dispatching activity work item", iid, task.Name, e.EventId)
			return nil, ctx.Err()
		}
		return nil, err
	}
	executor.logger.Debugf("%s: activity got result", key)
	if response == nil {
		return nil, errors.New("operation aborted")
	}

	var responseEvent *protos.HistoryEvent
	if failureDetails := response.GetFailureDetails(); failureDetails != nil {
		responseEvent = helpers.NewTaskFailedEvent(response.TaskId, response.FailureDetails)
	} else {
		responseEvent = helpers.NewTaskCompletedEvent(response.TaskId, response.Result)
	}

	return responseEvent, nil
}

// Shutdown implements Executor
func (g *grpcExecutor) Shutdown(ctx context.Context) error {
	// Iterate through all inflight work items and cancel them to unblock the goroutines waiting on ExecuteOrchestrator or ExecuteActivity
	g.inflightWorkItems.Range(func(key, value any) bool {
		cancel, ok := value.(context.CancelFunc)
		if ok {
			cancel()
		}
		g.inflightWorkItems.Delete(key)
		return true
	})

	return nil
}

// Hello implements protos.TaskHubSidecarServiceServer
func (grpcExecutor) Hello(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	return empty, nil
}

// GetWorkItems implements protos.TaskHubSidecarServiceServer
func (g *grpcExecutor) GetWorkItems(req *protos.GetWorkItemsRequest, stream protos.TaskHubSidecarService_GetWorkItemsServer) error {
	if md, ok := metadata.FromIncomingContext(stream.Context()); ok {
		g.logger.Infof("work item stream established by user-agent: %v", md.Get("user-agent"))
	}

	// There are some cases where the app may need to be notified when a client connects to fetch work items, like
	// for auto-starting the worker. The app also has an opportunity to set itself as unavailable by returning an error.
	callback := g.onWorkItemConnection
	if callback != nil {
		if err := callback(stream.Context()); err != nil {
			message := "unable to establish work item stream at this time: " + err.Error()
			g.logger.Warn(message)
			return status.Errorf(codes.Unavailable, message)
		}
	}

	defer func() {
		// Iterate through all inflight work items and cancel them to unblock the goroutines waiting on ExecuteOrchestrator or ExecuteActivity
		g.inflightWorkItems.Range(func(key, value any) bool {
			cancel, ok := value.(context.CancelFunc)
			if ok {
				cancel()
			}
			g.inflightWorkItems.Delete(key)
			return true
		})
	}()

	shutdownCtx, shutdownCancel := context.WithCancel(stream.Context())
	defer shutdownCancel()

	consumeCtx, cancel := context.WithCancel(shutdownCtx)
	defer cancel()

	go func() {
		<-g.streamShutdownChan
		shutdownCancel()
	}()

	// The worker client invokes this method, which streams back work-items as they arrive.
	err := g.backend.ConsumeWorkItems(consumeCtx, func(wi *GenericWorkItem) error {
		if err := stream.Send(wi); err != nil {
			g.logger.Errorf("encountered an error while sending work item: %v", err)
			return err
		}
		return nil
	})
	if err != nil {
		if stream.Context().Err() != nil {
			g.logger.Info("work item stream closed")
			return nil
		}
		if shutdownCtx.Err() != nil {
			return errShuttingDown
		}
	}
	return nil
}

// CompleteOrchestratorTask implements protos.TaskHubSidecarServiceServer
func (g *grpcExecutor) CompleteOrchestratorTask(ctx context.Context, res *protos.OrchestratorResponse) (*protos.CompleteTaskResponse, error) {
	err := g.backend.CompletePendingOrchestrator(ctx, res)
	if err != nil {
		return emptyCompleteTaskResponse, err
	}

	return emptyCompleteTaskResponse, nil
}

// CompleteActivityTask implements protos.TaskHubSidecarServiceServer
func (g *grpcExecutor) CompleteActivityTask(ctx context.Context, res *protos.ActivityResponse) (*protos.CompleteTaskResponse, error) {
	err := g.backend.CompletePendingActivity(ctx, res)
	if err != nil {
		return emptyCompleteTaskResponse, err
	}

	return emptyCompleteTaskResponse, nil
}

func getOrchestratorExecutionKey(iid string) string {
	return "wf:" + iid
}

func getActivityExecutionKey(iid string, taskID int32) string {
	return "activity:" + iid + "/" + strconv.FormatInt(int64(taskID), 10)
}

// CreateTaskHub implements protos.TaskHubSidecarServiceServer
func (grpcExecutor) CreateTaskHub(context.Context, *protos.CreateTaskHubRequest) (*protos.CreateTaskHubResponse, error) {
	return nil, errors.New("unimplemented")
}

// DeleteTaskHub implements protos.TaskHubSidecarServiceServer
func (grpcExecutor) DeleteTaskHub(context.Context, *protos.DeleteTaskHubRequest) (*protos.DeleteTaskHubResponse, error) {
	return nil, errors.New("unimplemented")
}

// GetInstance implements protos.TaskHubSidecarServiceServer
func (g *grpcExecutor) GetInstance(ctx context.Context, req *protos.GetInstanceRequest) (*protos.GetInstanceResponse, error) {
	metadata, err := g.backend.GetOrchestrationMetadata(ctx, api.InstanceID(req.InstanceId))
	if err != nil {
		return nil, err
	}
	if metadata == nil {
		return &protos.GetInstanceResponse{Exists: false}, nil
	}

	return createGetInstanceResponse(req, metadata), nil
}

// PurgeInstances implements protos.TaskHubSidecarServiceServer
func (g *grpcExecutor) PurgeInstances(ctx context.Context, req *protos.PurgeInstancesRequest) (*protos.PurgeInstancesResponse, error) {
	if req.GetPurgeInstanceFilter() != nil {
		return nil, errors.New("multi-instance purge is not unimplemented")
	}
	count, err := purgeOrchestrationState(ctx, g.backend, api.InstanceID(req.GetInstanceId()), req.Recursive)
	resp := &protos.PurgeInstancesResponse{DeletedInstanceCount: int32(count)}
	if err != nil {
		return resp, fmt.Errorf("failed to purge orchestration state: %w", err)
	}
	return resp, nil
}

// QueryInstances implements protos.TaskHubSidecarServiceServer
func (grpcExecutor) QueryInstances(context.Context, *protos.QueryInstancesRequest) (*protos.QueryInstancesResponse, error) {
	return nil, errors.New("unimplemented")
}

// RaiseEvent implements protos.TaskHubSidecarServiceServer
func (g *grpcExecutor) RaiseEvent(ctx context.Context, req *protos.RaiseEventRequest) (*protos.RaiseEventResponse, error) {
	e := helpers.NewEventRaisedEvent(req.Name, req.Input)
	if err := g.backend.AddNewOrchestrationEvent(ctx, api.InstanceID(req.InstanceId), e); err != nil {
		return nil, err
	}

	return &protos.RaiseEventResponse{}, nil
}

// StartInstance implements protos.TaskHubSidecarServiceServer
func (g *grpcExecutor) StartInstance(ctx context.Context, req *protos.CreateInstanceRequest) (*protos.CreateInstanceResponse, error) {
	instanceID := req.InstanceId
	ctx, span := helpers.StartNewCreateOrchestrationSpan(ctx, req.Name, req.Version.GetValue(), instanceID)
	defer span.End()

	e := helpers.NewExecutionStartedEvent(req.Name, instanceID, req.Input, nil, helpers.TraceContextFromSpan(span), req.ScheduledStartTimestamp)
	if err := g.backend.CreateOrchestrationInstance(ctx, e, WithOrchestrationIdReusePolicy(req.OrchestrationIdReusePolicy)); err != nil {
		return nil, err
	}

	return &protos.CreateInstanceResponse{InstanceId: instanceID}, nil
}

// TerminateInstance implements protos.TaskHubSidecarServiceServer
func (g *grpcExecutor) TerminateInstance(ctx context.Context, req *protos.TerminateRequest) (*protos.TerminateResponse, error) {
	e := helpers.NewExecutionTerminatedEvent(req.Output, req.Recursive)
	if err := g.backend.AddNewOrchestrationEvent(ctx, api.InstanceID(req.InstanceId), e); err != nil {
		return nil, fmt.Errorf("failed to submit termination request: %w", err)
	}
	return &protos.TerminateResponse{}, nil
}

// SuspendInstance implements protos.TaskHubSidecarServiceServer
func (g *grpcExecutor) SuspendInstance(ctx context.Context, req *protos.SuspendRequest) (*protos.SuspendResponse, error) {
	e := helpers.NewSuspendOrchestrationEvent(req.Reason.GetValue())
	if err := g.backend.AddNewOrchestrationEvent(ctx, api.InstanceID(req.InstanceId), e); err != nil {
		return nil, err
	}

	return &protos.SuspendResponse{}, nil
}

// ResumeInstance implements protos.TaskHubSidecarServiceServer
func (g *grpcExecutor) ResumeInstance(ctx context.Context, req *protos.ResumeRequest) (*protos.ResumeResponse, error) {
	e := helpers.NewResumeOrchestrationEvent(req.Reason.GetValue())
	if err := g.backend.AddNewOrchestrationEvent(ctx, api.InstanceID(req.InstanceId), e); err != nil {
		return nil, err
	}

	return &protos.ResumeResponse{}, nil
}

// WaitForInstanceCompletion implements protos.TaskHubSidecarServiceServer
func (g *grpcExecutor) WaitForInstanceCompletion(ctx context.Context, req *protos.GetInstanceRequest) (*protos.GetInstanceResponse, error) {
	return g.waitForInstance(ctx, req, func(m *api.OrchestrationMetadata) bool {
		return m.IsComplete()
	})
}

// WaitForInstanceStart implements protos.TaskHubSidecarServiceServer
func (g *grpcExecutor) WaitForInstanceStart(ctx context.Context, req *protos.GetInstanceRequest) (*protos.GetInstanceResponse, error) {
	return g.waitForInstance(ctx, req, func(m *api.OrchestrationMetadata) bool {
		return m.RuntimeStatus != protos.OrchestrationStatus_ORCHESTRATION_STATUS_PENDING
	})
}

func (g *grpcExecutor) waitForInstance(ctx context.Context, req *protos.GetInstanceRequest, condition func(*api.OrchestrationMetadata) bool) (*protos.GetInstanceResponse, error) {
	iid := api.InstanceID(req.InstanceId)

	var b backoff.BackOff = &backoff.ExponentialBackOff{
		InitialInterval:     1 * time.Millisecond,
		MaxInterval:         3 * time.Second,
		Multiplier:          1.5,
		RandomizationFactor: 0.5,
		Stop:                backoff.Stop,
		Clock:               backoff.SystemClock,
	}
	b = backoff.WithContext(b, ctx)
	b.Reset()

loop:
	for {
		t := time.NewTimer(b.NextBackOff())
		select {
		case <-ctx.Done():
			if !t.Stop() {
				<-t.C
			}
			break loop

		case <-t.C:
			metadata, err := g.backend.GetOrchestrationMetadata(ctx, iid)
			if err != nil {
				return nil, err
			}
			if metadata == nil {
				return &protos.GetInstanceResponse{Exists: false}, nil
			}
			if condition(metadata) {
				return createGetInstanceResponse(req, metadata), nil
			}
		}
	}

	return nil, status.Errorf(codes.Canceled, "instance hasn't completed")
}

// mustEmbedUnimplementedTaskHubSidecarServiceServer implements protos.TaskHubSidecarServiceServer
func (grpcExecutor) mustEmbedUnimplementedTaskHubSidecarServiceServer() {
}

func createGetInstanceResponse(req *protos.GetInstanceRequest, metadata *api.OrchestrationMetadata) *protos.GetInstanceResponse {
	state := &protos.OrchestrationState{
		InstanceId:           req.InstanceId,
		Name:                 metadata.Name,
		OrchestrationStatus:  metadata.RuntimeStatus,
		CreatedTimestamp:     timestamppb.New(metadata.CreatedAt),
		LastUpdatedTimestamp: timestamppb.New(metadata.LastUpdatedAt),
	}

	if req.GetInputsAndOutputs {
		state.Input = wrapperspb.String(metadata.SerializedInput)
		state.CustomStatus = wrapperspb.String(metadata.SerializedCustomStatus)
		state.Output = wrapperspb.String(metadata.SerializedOutput)
		state.FailureDetails = metadata.FailureDetails
	}

	return &protos.GetInstanceResponse{Exists: true, OrchestrationState: state}
}
