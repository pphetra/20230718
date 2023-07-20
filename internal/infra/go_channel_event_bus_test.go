package infra_test

import (
	"context"
	"testing"
	"time"

	"taejai/internal/infra"
	member_domain "taejai/internal/member/domain"
	member_domain_event "taejai/internal/member/domain/event"
	shared_app "taejai/internal/shared/app"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGoChannelEventBus_Publish(t *testing.T) {
	bus := infra.NewGoChannelEventBus()

	received := make(chan member_domain_event.IndividualMemberRegisteredEvent)

	sub, err := bus.Pubsub.Subscribe(context.Background(), "events")
	assert.NoError(t, err)

	go func() {
		msg := <-sub
		domain_event_name := msg.Metadata.Get("event_name")
		assert.Equal(t, member_domain_event.IndividualMemberRegisteredEventName, domain_event_name)

		imrEvent, err := member_domain_event.NewIndividualMemberRegisteredEventFromJsonBytes(msg.Payload)
		assert.NoError(t, err)

		received <- imrEvent
	}()

	bus.Publish(member_domain_event.NewIndividualMemberRegisteredEvent(member_domain.MemberId(1)))

	select {
	case event := <-received:
		assert.Equal(t, member_domain_event.IndividualMemberRegisteredEventName, event.GetName())
		assert.Equal(t, member_domain.MemberId(1), event.MemberId)
	case <-time.After(1 * time.Second):
		t.Fatal("timeout")
	}

	// clean up
	bus.Pubsub.Close()
}

type TestEvent struct {
}

func (e TestEvent) GetName() string {
	return "test_event"
}

func (e TestEvent) GetPayload() []byte {
	return []byte("test")
}

func TestGoChannelEventBus_Handler(t *testing.T) {
	t.Parallel()
	mockUOW := shared_app.MockUnitOfWork{}
	bus := infra.NewGoChannelEventBus()
	commandDispatcher := shared_app.NewCommandDispatcher(
		&mockUOW,
		bus.Publish,
	)
	bus.CommandDipatcher = &commandDispatcher

	mockEventHandler := shared_app.MockEventHandler{}
	mockEventHandler.On("Handle", mock.Anything, TestEvent{}).Return(nil)
	mockEventHandler.On("ParseEvent", mock.Anything).Return(TestEvent{}, nil)

	bus.RegisterHandler("test_event", &mockEventHandler)

	bus.Start()
	defer bus.Stop()

	bus.Publish(TestEvent{})

	time.Sleep(1 * time.Second)

	mockEventHandler.AssertCalled(t, "Handle", mock.Anything, mock.Anything)
	mockEventHandler.AssertCalled(t, "ParseEvent", mock.Anything)
}
