package es

type CommandHandler interface {
	ProcessCommand(command Command) error
}

type commandHandler struct {
	Events EventStore
}

func NewCommandHandler(events EventStore) CommandHandler {
	return &commandHandler{
		Events: events,
	}
}

func (handler *commandHandler) ProcessCommand(command Command) error {
	newEvents, err := command.Execute(handler.Events)
	if err != nil {
		return err
	}
	return handler.Events.Append(newEvents...)
}

type EventFetcher interface {
	Fetch(filter EventFilter) ([]Event, error)
}

// EventStore abstracts the event store.
type EventStore interface {
	Append(event ...Event) error
	Fetch(filter EventFilter) ([]Event, error)
}

var _ EventStore = &Events{}

type Command interface {
	Execute(EventFetcher) ([]Event, error)
}
