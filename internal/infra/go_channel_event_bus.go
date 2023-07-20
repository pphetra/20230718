package infra

import (
	"taejai/internal/shared/value_object"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

type GoChannelEventBus struct {
	Pubsub *gochannel.GoChannel
}

func NewGoChannelEventBus() GoChannelEventBus {
	pubSub := gochannel.NewGoChannel(
		gochannel.Config{},
		watermill.NewStdLogger(false, false),
	)
	return GoChannelEventBus{
		Pubsub: pubSub,
	}
}

func (b GoChannelEventBus) Publish(event value_object.DomainEvent) error {
	msg := message.NewMessage(watermill.NewUUID(), event.GetPayload())
	msg.Metadata.Set("event_name", event.GetName())
	return b.Pubsub.Publish(event.GetName(), msg)
}
