package shared_app

import (
	"taejai/internal/shared/value_object"

	"github.com/stretchr/testify/mock"
)

type MockUnitOfWork struct {
	mock.Mock
}

func (m *MockUnitOfWork) GetRepository(key string) interface{} {
	args := m.Called(key)
	return args.Get(0)
}

func (m *MockUnitOfWork) Publish(event value_object.DomainEvent) error {
	args := m.Called(event)
	return args.Error(0)
}

func (m *MockUnitOfWork) DoInTransaction(fn TxFunc) (interface{}, error) {
	return fn(
		m,
		m.Publish,
	)
}
