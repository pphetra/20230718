package member_app_handler

import (
	member_app "taejai/internal/member/app"
	member_domain_event "taejai/internal/member/domain/event"
	shared_app "taejai/internal/shared/app"
	shared_domain "taejai/internal/shared/domain"
)

type IndividualMemberRegisteredHandler struct {
}

func (h IndividualMemberRegisteredHandler) Handle(dispatcher *shared_app.CommandDispatcher, event shared_domain.DomainEvent) error {
	// TODO
	// dispatcher.Dispatch(sendGreetingEmailCommand)
	imrEvent := event.(member_domain_event.IndividualMemberRegisteredEvent)

	// TODO
	// new sendGreetingEmailCommand
	command := member_app.SendGreetingCommand{
		MemberId: imrEvent.MemberId,
	}

	dispatcher.Execute(command)

	return nil
}

func (h IndividualMemberRegisteredHandler) ParseEvent(payload []byte) (shared_domain.DomainEvent, error) {
	// parse payload to event
	return member_domain_event.NewIndividualMemberRegisteredEventFromJsonBytes(payload)
}
