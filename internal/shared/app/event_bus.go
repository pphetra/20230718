package shared_app

import "taejai/internal/shared/value_object"

type PublishEvent func(event value_object.DomainEvent) error

type EventBus interface {
	Publish(event value_object.DomainEvent) error
	RegisterHandler(eventName string, handler EventHandler) error
}

type EventHandler interface {
	Handle(dispatcher *CommandDispatcher, event value_object.DomainEvent) error
	ParseEvent(payload []byte) (value_object.DomainEvent, error)
}
