package member_app_test

import (
	"errors"
	member_app "taejai/internal/member/app"
	member_domain "taejai/internal/member/domain"
	shared_app "taejai/internal/shared/app"
	"taejai/internal/shared/value_object"
	"testing"

	member_domain_event "taejai/internal/member/domain/event"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterUseCase_Register_success_with_published_event(t *testing.T) {
	mockUOW := &shared_app.MockUnitOfWork{}
	mockRepository := &member_domain.MockMemberRepository{}

	mockUOW.On("GetRepository", "member").Return(mockRepository)
	mockUOW.On("Publish",
		member_domain_event.IndividualMemberRegisteredEvent{MemberId: member_domain.MemberId(1)},
	).Return(nil)

	address, err := value_object.NewAddress(
		"123/456",
		"Bangkok",
		"10110",
	)
	assert.NoError(t, err)

	member, err := member_domain.NewIndividualMember(
		"John",
		"Doe",
		address,
	)
	assert.NoError(t, err)

	mockRepository.On("Create", &member).Return(member_domain.MemberId(1), nil)

	id, err := member_app.RegisterIndividualMember(
		mockUOW,
		member.Name1,
		member.Name2,
		address.Line1,
		address.Line2,
		address.PostalCode,
	)

	assert.NoError(t, err)
	assert.Equal(t, member_domain.MemberId(1), id)
	mockUOW.AssertExpectations(t)

}

func TestRegisterUseCase_Register_rollback_without_published_event(t *testing.T) {
	mockUOW := &shared_app.MockUnitOfWork{}
	mockRepository := &member_domain.MockMemberRepository{}

	mockUOW.On("GetRepository", "member").Return(mockRepository)

	mockRepository.On("Create", mock.Anything).Return(member_domain.MemberId(0), errors.New("error"))

	_, err := member_app.RegisterIndividualMember(
		mockUOW,
		"John",
		"Doe",
		"123/456",
		"Bangkok",
		"10110",
	)

	assert.Error(t, err)
	mockUOW.AssertExpectations(t)

}
