package infra_test

import (
	"context"
	"testing"
	"time"

	"taejai/internal/infra"
	member_domain "taejai/internal/member/domain"
	member_domain_event "taejai/internal/member/domain/event"

	"github.com/stretchr/testify/assert"
)

func TestGoChannelEventBus_Publish(t *testing.T) {
	bus := infra.NewGoChannelEventBus()

	received := make(chan member_domain_event.IndividualMemberRegisteredEvent)

	sub, err := bus.Pubsub.Subscribe(context.Background(), member_domain_event.IndividualMemberRegisteredEventName)
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
