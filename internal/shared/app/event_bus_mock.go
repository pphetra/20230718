package shared_app

import (
	"taejai/internal/shared/value_object"

	"github.com/stretchr/testify/mock"
)

type MockEventBus struct {
	mock.Mock
}

func (m *MockEventBus) Publish(event value_object.DomainEvent) error {
	args := m.Called(event)
	return args.Error(0)
}

func (m *MockEventBus) RegisterHandler(eventName string, handler EventHandler) error {
	args := m.Called(eventName, handler)
	return args.Error(0)
}

type MockEventHandler struct {
	mock.Mock
}

func (m *MockEventHandler) Handle(dispatcher *CommandDispatcher, event value_object.DomainEvent) error {
	args := m.Called(dispatcher, event)
	return args.Error(0)
}

func (m *MockEventHandler) ParseEvent(payload []byte) (value_object.DomainEvent, error) {
	args := m.Called(payload)
	return args.Get(0).(value_object.DomainEvent), args.Error(1)
}
