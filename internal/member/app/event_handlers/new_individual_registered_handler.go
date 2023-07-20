package member_app_handler

import (
	member_app_commands "taejai/internal/member/app/commands"
	member_domain_events "taejai/internal/member/domain/events"
	shared_app "taejai/internal/shared/app"
	shared_domain "taejai/internal/shared/domain"
)

type IndividualMemberRegisteredHandler struct {
}

func (h IndividualMemberRegisteredHandler) Handle(dispatcher *shared_app.CommandDispatcher, event shared_domain.DomainEvent) error {
	// TODO
	// dispatcher.Dispatch(sendGreetingEmailCommand)
	imrEvent := event.(member_domain_events.IndividualMemberRegisteredEvent)

	// TODO
	// new sendGreetingEmailCommand
	command := member_app_commands.SendGreetingCommand{
		MemberId: imrEvent.MemberId,
	}

	dispatcher.Execute(command)

	return nil
}

func (h IndividualMemberRegisteredHandler) ParseEvent(payload []byte) (shared_domain.DomainEvent, error) {
	// parse payload to event
	return member_domain_events.NewIndividualMemberRegisteredEventFromJsonBytes(payload)
}
