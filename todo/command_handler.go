package todo

import "github.com/GodsBoss/es-todo/es"

type CommandHandler struct {
	Events EventStore
}

func (handler *CommandHandler) ProcessTaskCommand(command Command) error {
	newEvents, err := command.Execute(handler.Events)
	if err != nil {
		return err
	}
	return handler.Events.Append(newEvents...)
}

func loadTask(fetcher EventFetcher, id string) (*Task, error) {
	pastEvents, err := fetcher.Fetch(es.ByAggregateID(id))
	if err != nil {
		return nil, err
	}
	task := &Task{}
	for i := range pastEvents {
		task.apply(pastEvents[i])
	}
	return task, nil
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
