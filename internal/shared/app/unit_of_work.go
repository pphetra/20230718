package shared_app

import "taejai/internal/shared/value_object"

type PublishEvent func(value_object.DomainEvent) error

type TxFunc func(UnitOfWorkRepositoryStore, PublishEvent) (interface{}, error)

type UnitOfWorkRepositoryStore interface {
	GetRepository(key string) interface{}
}

type UnitOfWork interface {
	DoInTransaction(f TxFunc) (interface{}, error)
}
