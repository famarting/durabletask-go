package backend

import (
	"encoding/json"
	"testing"

	"github.com/microsoft/durabletask-go/api"
	"github.com/microsoft/durabletask-go/internal/helpers"
	"github.com/microsoft/durabletask-go/internal/protos"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestOrchestrationRuntimeStateSerialization(t *testing.T) {

	runtimeState := ConvertToSerializableRuntimeState(NewOrchestrationRuntimeState(api.InstanceID("foo"), []*protos.HistoryEvent{
		helpers.NewExecutionStartedEvent("test", "foo", nil, nil, nil, timestamppb.Now()),
		helpers.NewTaskScheduledEvent(0, "testtask", wrapperspb.String("foo"), nil, nil),
	}))

	raw, err := json.Marshal(runtimeState)
	require.NoError(t, err)

	var out *SerializableOrchestrationRuntimeState
	err = json.Unmarshal(raw, &out)
	require.NoError(t, err)

	raw2, err := json.Marshal(out)
	require.NoError(t, err)

	require.Equal(t, raw, raw2)
}

func TestSerializeHistoryEvent(t *testing.T) {

	e := &internalHistoryEvent{
		HistoryEvent: helpers.NewEventRaisedEvent("aa", wrapperspb.String("foo")),
	}

	raw, err := json.Marshal(e)
	require.NoError(t, err)

	var out *internalHistoryEvent
	err = json.Unmarshal(raw, &out)
	require.NoError(t, err)

	require.NotNil(t, out.GetEventRaised())
}
