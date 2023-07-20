package member_domain

import "github.com/stretchr/testify/mock"

type MockMemberRepository struct {
	mock.Mock
}

func (m *MockMemberRepository) GetById(id MemberId) (*Member, error) {
	args := m.Called(id)
	return args.Get(0).(*Member), args.Error(1)
}

func (m *MockMemberRepository) FindByName(name string) ([]*Member, error) {
	args := m.Called(name)
	return args.Get(0).([]*Member), args.Error(1)
}

func (m *MockMemberRepository) Create(member *Member) (MemberId, error) {
	args := m.Called(member)
	return args.Get(0).(MemberId), args.Error(1)
}

func (m *MockMemberRepository) Update(member *Member) error {
	args := m.Called(member)
	return args.Error(0)
}

func (m *MockMemberRepository) Delete(member *Member) error {
	args := m.Called(member)
	return args.Error(0)
}
