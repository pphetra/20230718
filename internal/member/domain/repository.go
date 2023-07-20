package member_domain

type MemberRepository interface {
	GetById(id MemberId) (*Member, error)
	FindByName(name string) ([]*Member, error)
	Create(member *Member) (MemberId, error)
	Update(member *Member) error
	Delete(member *Member) error
}
