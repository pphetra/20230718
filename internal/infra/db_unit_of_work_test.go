package infra_test

import (
	"errors"
	"taejai/internal/infra"
	member_domain "taejai/internal/member/domain"
	shared_app "taejai/internal/shared/app"
	"taejai/internal/shared/value_object"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// Feature: PostgresUnitOfWork_DoInTransaction
// Scenario: Successfully create a new member
//
//	Given a database connection
//	And a PostgresUnitOfWork instance
//	And a member's information
//	When DoInTransaction is called with a function that creates a new member
//	Then the transaction should be committed
//	And the created member should be returned
func TestPostgresUnitOfWork_DoInTransaction(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	eventBus := infra.NewGoChannelEventBus()
	uow := infra.NewDBUnitOfWork(db, infra.WithEventBus(eventBus))

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO members").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	result, err := uow.DoInTransaction(func(store shared_app.UnitOfWorkStore, publish shared_app.PublishEvent) (interface{}, error) {
		memberRepo := store.GetRepository("member").(member_domain.MemberRepository)
		address, err := value_object.NewAddress(
			"123/456",
			"Bangkok",
			"10110",
		)
		if err != nil {
			return nil, err
		}

		member, err := member_domain.NewIndividualMember("John", "Doe", address)
		if err != nil {
			return nil, err
		}

		_, err = memberRepo.Create(&member)
		if err != nil {
			return nil, err
		}
		return member, nil
	})

	assert.NoError(t, err)
	assert.NotNil(t, result)

}

// Scenario: Failed to create a new member
// Given a database connection
// And a PostgresUnitOfWork instance
// And a member's information
// When DoInTransaction is called with a function that fails to create a new member
// Then the transaction should be rolled back
// And an error should be returned
func TestPostgresUnitOfWork_DoInTransaction_Rollback(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	uow := infra.NewDBUnitOfWork(db)

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO members").WillReturnError(errors.New("failed to insert member"))
	mock.ExpectRollback()

	_, err = uow.DoInTransaction(func(store shared_app.UnitOfWorkStore, publish shared_app.PublishEvent) (interface{}, error) {
		memberRepo := store.GetRepository("member").(member_domain.MemberRepository)
		address, err := value_object.NewAddress(
			"123/456",
			"Bangkok",
			"10110",
		)
		if err != nil {
			return nil, err
		}

		member, err := member_domain.NewIndividualMember("John", "Doe", address)
		if err != nil {
			return nil, err
		}
		_, err = memberRepo.Create(&member)
		if err != nil {
			return nil, err
		}
		return member, nil
	})

	assert.Error(t, err)
}
