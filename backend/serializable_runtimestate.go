package backend

import (
	"time"

	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/microsoft/durabletask-go/api"

	"google.golang.org/protobuf/encoding/protojson"
)

type SerializableOrchestrationRuntimeState struct {
	InstanceID      api.InstanceID
	NewEvents       []*internalHistoryEvent
	OldEvents       []*internalHistoryEvent
	PendingTasks    []*internalHistoryEvent
	PendingTimers   []*internalHistoryEvent
	PendingMessages []internalOrchestratorMessage

	StartEvent      *ExecutionStartedEvent
	CompletedEvent  *ExecutionCompletedEvent
	CreatedTime     time.Time
	LastUpdatedTime time.Time
	CompletedTime   time.Time
	ContinuedAsNew  bool
	IsSuspended     bool

	CustomStatus *wrapperspb.StringValue
}

func ConvertToSerializableRuntimeState(in *OrchestrationRuntimeState) *SerializableOrchestrationRuntimeState {
	out := &SerializableOrchestrationRuntimeState{
		InstanceID:      in.InstanceID(),
		NewEvents:       convertToInternalEvents(in.NewEvents()),
		OldEvents:       convertToInternalEvents(in.OldEvents()),
		PendingTasks:    convertToInternalEvents(in.PendingTasks()),
		PendingTimers:   convertToInternalEvents(in.PendingTimers()),
		PendingMessages: convertToInternalOrchestratorMessage(in.PendingMessages()),

		StartEvent:      in.startEvent,
		CompletedEvent:  in.completedEvent,
		CreatedTime:     in.createdTime,
		LastUpdatedTime: in.lastUpdatedTime,
		CompletedTime:   in.completedTime,
		ContinuedAsNew:  in.continuedAsNew,
		IsSuspended:     in.isSuspended,

		CustomStatus: in.CustomStatus,
	}
	return out
}

func ConvertFromSerializableRuntimeState(in *SerializableOrchestrationRuntimeState) *OrchestrationRuntimeState {
	out := &OrchestrationRuntimeState{
		instanceID:      in.InstanceID,
		newEvents:       convertFromInternalEvents(in.NewEvents),
		oldEvents:       convertFromInternalEvents(in.OldEvents),
		pendingTasks:    convertFromInternalEvents(in.PendingTasks),
		pendingTimers:   convertFromInternalEvents(in.PendingTimers),
		pendingMessages: convertFromInternalOrchestratorMessage(in.PendingMessages),

		startEvent:      in.StartEvent,
		completedEvent:  in.CompletedEvent,
		createdTime:     in.CreatedTime,
		lastUpdatedTime: in.LastUpdatedTime,
		completedTime:   in.CompletedTime,
		continuedAsNew:  in.ContinuedAsNew,
		isSuspended:     in.IsSuspended,

		CustomStatus: in.CustomStatus,
	}
	return out
}

type internalOrchestratorMessage struct {
	HistoryEvent     *internalHistoryEvent
	TargetInstanceID string
}

type internalHistoryEvent struct {
	*HistoryEvent
}

func (h *internalHistoryEvent) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(h.HistoryEvent)
}
func (h *internalHistoryEvent) UnmarshalJSON(data []byte) (err error) {
	e := &HistoryEvent{}
	err = protojson.Unmarshal(data, e)
	if err != nil {
		return err
	}
	h.HistoryEvent = e
	return nil
}

func convertFromInternalEvents(in []*internalHistoryEvent) []*HistoryEvent {
	out := []*HistoryEvent{}
	for _, i := range in {
		out = append(out, i.HistoryEvent)
	}
	return out
}

func convertToInternalEvents(in []*HistoryEvent) []*internalHistoryEvent {
	out := []*internalHistoryEvent{}
	for _, i := range in {
		out = append(out, &internalHistoryEvent{
			HistoryEvent: i,
		})
	}
	return out
}

func convertFromInternalOrchestratorMessage(in []internalOrchestratorMessage) []OrchestratorMessage {
	out := []OrchestratorMessage{}
	for _, i := range in {
		out = append(out, OrchestratorMessage{
			HistoryEvent:     i.HistoryEvent.HistoryEvent,
			TargetInstanceID: i.TargetInstanceID,
		})
	}
	return out
}

func convertToInternalOrchestratorMessage(in []OrchestratorMessage) []internalOrchestratorMessage {
	out := []internalOrchestratorMessage{}
	for _, i := range in {
		out = append(out, internalOrchestratorMessage{
			HistoryEvent:     &internalHistoryEvent{i.HistoryEvent},
			TargetInstanceID: i.TargetInstanceID,
		})
	}
	return out
}
