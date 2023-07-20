package infra

import (
	"context"
	"sync"
	"taejai/internal/shared/value_object"

	shared_app "taejai/internal/shared/app"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

type GoChannelEventBus struct {
	Pubsub           *gochannel.GoChannel
	handlers         map[string][]shared_app.EventHandler
	CommandDipatcher *shared_app.CommandDispatcher
	stopChan         chan struct{}
	wg               sync.WaitGroup
}

type GoChannelEventBusOption func(*GoChannelEventBus)

func WithCommandDispatcher(dispatcher *shared_app.CommandDispatcher) GoChannelEventBusOption {
	return func(bus *GoChannelEventBus) {
		bus.CommandDipatcher = dispatcher
	}
}

func NewGoChannelEventBus(options ...GoChannelEventBusOption) *GoChannelEventBus {
	pubSub := gochannel.NewGoChannel(
		gochannel.Config{},
		watermill.NewStdLogger(false, false),
	)
	instance := GoChannelEventBus{
		Pubsub:   pubSub,
		handlers: make(map[string][]shared_app.EventHandler),
		stopChan: make(chan struct{}),
	}

	for _, option := range options {
		option(&instance)
	}

	return &instance
}

func (b *GoChannelEventBus) Publish(event value_object.DomainEvent) error {
	msg := message.NewMessage(watermill.NewUUID(), event.GetPayload())
	msg.Metadata.Set("event_name", event.GetName())
	// all events go to the same topic
	// handler will use MetaData["event_name"] to determine which handlers will handle the event
	return b.Pubsub.Publish("events", msg)
}

func (b *GoChannelEventBus) RegisterHandler(eventName string, handler shared_app.EventHandler) error {
	// check handlers[eventName] is exists
	handlers, ok := b.handlers[eventName]

	// if not exists, create handlers[eventName] = []EventHandler{}
	if !ok {
		b.handlers[eventName] = []shared_app.EventHandler{}
	}

	// append handler to handlers[eventName]
	b.handlers[eventName] = append(handlers, handler)

	return nil
}

func (b *GoChannelEventBus) Start() error {
	// subscribe to topic "events"
	// when message is received, call b.handleMessage
	messages, err := b.Pubsub.Subscribe(context.Background(), "events")

	if err != nil {
		return err
	}

	go func() {
		for event := range messages {
			b.handleMessage(event)
		}
	}()

	return nil
}

func (b *GoChannelEventBus) Stop() {
	close(b.stopChan)
	b.wg.Wait()
}

func (b *GoChannelEventBus) handleMessage(message *message.Message) {
	// get event_name from event.MetaData
	eventName := message.Metadata.Get("event_name")

	// get handlers[eventName]
	handlers, ok := b.handlers[eventName]
	if !ok {
		return
	}

	// call handler.Handle(event)
	for _, handler := range handlers {
		event, err := handler.ParseEvent(message.Payload)
		if err != nil {
			// TODO log error
		} else {
			handler.Handle(b.CommandDipatcher, event)
		}
	}
}
