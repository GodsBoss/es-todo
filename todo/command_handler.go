package todo

import (
	"github.com/GodsBoss/es-todo/es"
)

type CommandHandler struct {
	Events EventStore
}

func (handler *CommandHandler) ProcessCommand(command Command) error {
	newEvents, err := command.Execute(handler.Events)
	if err != nil {
		return err
	}
	return handler.Events.Append(newEvents...)
}

type EventFetcher interface {
	Fetch(filter es.EventFilter) ([]es.Event, error)
}

// EventStore abstracts the event store.
type EventStore interface {
	Append(event ...es.Event) error
	Fetch(filter es.EventFilter) ([]es.Event, error)
}

var _ EventStore = &es.Events{}

type Command interface {
	Execute(EventFetcher) ([]es.Event, error)
}
