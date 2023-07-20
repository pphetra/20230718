package shared_app

import (
	shared_domain "taejai/internal/shared/domain"

	"github.com/stretchr/testify/mock"
)

type MockUnitOfWork struct {
	mock.Mock
}

func (m *MockUnitOfWork) GetRepository(key string) interface{} {
	args := m.Called(key)
	return args.Get(0)
}

func (m *MockUnitOfWork) Publish(event shared_domain.DomainEvent) error {
	args := m.Called(event)
	return args.Error(0)
}

func (m *MockUnitOfWork) DoInTransaction(fn TxFunc) (interface{}, error) {
	return fn(
		m,
		m.Publish,
	)
}
