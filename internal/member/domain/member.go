package member_domain

type MemberType int

const (
	Individual MemberType = iota
	Organization
)

type MemberId int64

type Member struct {
	Id      MemberId
	Name1   string
	Name2   string
	Type    MemberType
	Address Address
}

func NewIndividualMember(
	firstName string,
	lastName string,
	address Address,
) (Member, error) {
	return Member{
		Name1:   firstName,
		Name2:   lastName,
		Type:    Individual,
		Address: address,
	}, nil
}
