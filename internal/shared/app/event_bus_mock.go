package shared_app

import (
	shared_domain "taejai/internal/shared/domain"

	"github.com/stretchr/testify/mock"
)

type MockEventBus struct {
	mock.Mock
}

func (m *MockEventBus) Publish(event shared_domain.DomainEvent) error {
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

func (m *MockEventHandler) Handle(dispatcher *CommandDispatcher, event shared_domain.DomainEvent) error {
	args := m.Called(dispatcher, event)
	return args.Error(0)
}

func (m *MockEventHandler) ParseEvent(payload []byte) (shared_domain.DomainEvent, error) {
	args := m.Called(payload)
	return args.Get(0).(shared_domain.DomainEvent), args.Error(1)
}
