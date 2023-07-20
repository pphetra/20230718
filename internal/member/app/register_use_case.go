package member_app

import (
	member_domain "taejai/internal/member/domain"
	member_domain_events "taejai/internal/member/domain/events"
	shared_app "taejai/internal/shared/app"
)

func RegisterIndividualMember(
	unitOfWork shared_app.UnitOfWork,
	firstName string,
	lastName string,
	addressLine1 string,
	addressLine2 string,
	addressPostalCode string,
) (member_domain.MemberId, error) {
	id, err := unitOfWork.DoInTransaction(func(store shared_app.UnitOfWorkStore, publish shared_app.PublishEvent) (interface{}, error) {
		// why not memberRepo := store.GetMemberRepository()?
		memberRepository := store.GetRepository("member").(member_domain.MemberRepository)

		// what is the error handling strategy here?
		address, err := member_domain.NewAddress(
			addressLine1,
			addressLine2,
			addressPostalCode,
		)
		if err != nil {
			return 0, err
		}

		member, err := member_domain.NewIndividualMember(
			firstName,
			lastName,
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
		publish(member_domain_events.NewIndividualMemberRegisteredEvent(id))

		return id, err
	})
	if err != nil {
		return 0, err
	}
	return id.(member_domain.MemberId), nil
}
