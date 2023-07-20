package shared_app

import (
	shared_domain "taejai/internal/shared/domain"
)

type PublishEvent func(event shared_domain.DomainEvent) error

type EventBus interface {
	Publish(event shared_domain.DomainEvent) error
	RegisterHandler(eventName string, handler EventHandler) error
}

type EventHandler interface {
	Handle(dispatcher *CommandDispatcher, event shared_domain.DomainEvent) error
	ParseEvent(payload []byte) (shared_domain.DomainEvent, error)
}
