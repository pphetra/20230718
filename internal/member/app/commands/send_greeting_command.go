package member_app

import (
	"fmt"
	member_domain "taejai/internal/member/domain"
	shared_app "taejai/internal/shared/app"
)

type SendGreetingCommand struct {
	MemberId member_domain.MemberId
}

func (c SendGreetingCommand) GetName() string {
	return "send_greeting"
}

func (c SendGreetingCommand) Execute(store shared_app.UnitOfWorkStore, publish shared_app.PublishEvent) (interface{}, error) {
	repo := store.GetRepository("member").(member_domain.MemberRepository)
	member, err := repo.GetById(c.MemberId)
	if err != nil {
		return nil, err
	}
	fmt.Println("send greeting to", member)

	// generate message based on member information

	// send mail to member
	// how to access mail service?

	// mark member as greeted
	// save member

	return nil, nil
}
