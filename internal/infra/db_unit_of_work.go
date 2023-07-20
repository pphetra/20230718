package infra

import (
	"database/sql"
	member_infra "taejai/internal/member/infra"
	shared_app "taejai/internal/shared/app"
	"taejai/internal/shared/value_object"
)

type DBUnitOfWorkStore struct {
	tx *sql.Tx
}

func (s *DBUnitOfWorkStore) GetRepository(key string) interface{} {
	switch key {
	case "member":
		repo := member_infra.NewMemberRepository(s.tx)
		return repo
	default:
		return nil
	}
}

type DBUnitOfWork struct {
	db       *sql.DB
	eventBus shared_app.EventBus
}

type DBUnitOfWorkOption func(*DBUnitOfWork)

func WithEventBus(eventBus shared_app.EventBus) func(*DBUnitOfWork) {
	return func(u *DBUnitOfWork) {
		u.eventBus = eventBus
	}
}

func NewDBUnitOfWork(db *sql.DB, options ...DBUnitOfWorkOption) *DBUnitOfWork {
	uow := &DBUnitOfWork{db: db}

	for _, option := range options {
		option(uow)
	}

	return uow
}

func (u *DBUnitOfWork) DoInTransaction(txFunc shared_app.TxFunc) (interface{}, error) {
	tx, err := u.db.Begin()
	if err != nil {
		return nil, err
	}

	store := &DBUnitOfWorkStore{tx: tx}

	events := []value_object.DomainEvent{}
	publishFunc := func(event value_object.DomainEvent) error {
		events = append(events, event)
		return nil
	}

	result, err := txFunc(store, publishFunc)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	for _, event := range events {
		err = u.eventBus.Publish(event)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
