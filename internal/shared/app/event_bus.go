package shared_app

import "taejai/internal/shared/value_object"

type EventBus interface {
	Publish(event value_object.DomainEvent) error
}
