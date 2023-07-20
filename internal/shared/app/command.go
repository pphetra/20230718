package shared_app

type Command interface {
	GetName() string
	Execute(UnitOfWorkStore, PublishEvent) (interface{}, error)
}

type CommandDispatcher struct {
	uow     UnitOfWork
	publish PublishEvent
}

func NewCommandDispatcher(uow UnitOfWork, publish PublishEvent) CommandDispatcher {
	return CommandDispatcher{
		uow:     uow,
		publish: publish,
	}
}

// command will execute in a transaction
// if command is successful, it will commit the transaction
// ? how to inject external service into command? eg. mail service
func (c CommandDispatcher) Execute(command Command) (interface{}, error) {
	return c.uow.DoInTransaction(func(store UnitOfWorkStore, publish PublishEvent) (interface{}, error) {
		return command.Execute(store, publish)
	})
}
