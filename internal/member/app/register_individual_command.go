package member_app

import (
	member_domain "taejai/internal/member/domain"
	member_domain_event "taejai/internal/member/domain/event"
	shared_app "taejai/internal/shared/app"
)

type RegisterIndividualCommand struct {
	FirstName         string
	LastName          string
	addressLine1      string
	addressLine2      string
	addressPostalCode string
}

func (c RegisterIndividualCommand) GetName() string {
	return "register_individual"
}

func (c RegisterIndividualCommand) Execute(store shared_app.UnitOfWorkStore, publish shared_app.PublishEvent) (interface{}, error) {
	// why not memberRepo := store.GetMemberRepository()?
	memberRepository := store.GetRepository("member").(member_domain.MemberRepository)

	// what is the error handling strategy here?
	address, err := member_domain.NewAddress(
		c.addressLine1,
		c.addressLine2,
		c.addressPostalCode,
	)
	if err != nil {
		return 0, err
	}

	member, err := member_domain.NewIndividualMember(
		c.FirstName,
		c.LastName,
		address,
	)
	if err != nil {
		return 0, err
	}
	// should we pass member by value or by reference?
	id, err := memberRepository.Create(&member)
	if err != nil {
		return 0, err
	}

	// publish RegisteredEvent
	publish(member_domain_event.NewIndividualMemberRegisteredEvent(id))

	return id, err
}
