package shared_app

type TxFunc func(UnitOfWorkStore, PublishEvent) (interface{}, error)

// hold all repositories and external services
type UnitOfWorkStore interface {
	GetRepository(key string) interface{}
}

type UnitOfWork interface {
	DoInTransaction(f TxFunc) (interface{}, error)
}
