package todo

import "github.com/GodsBoss/es-todo/es"

type CommandHandler struct {
	Events EventStore
}

func (handler *CommandHandler) ProcessTaskCommand(command TaskCommand) error {
	pastEvents, err := handler.Events.Fetch(es.ByAggregateID(command.AggregateID()))
	if err != nil {
		return err
	}
	task := &Task{}
	for i := range pastEvents {
		task.apply(pastEvents[i])
	}
	newEvents, err := task.process(command)
	if err != nil {
		return err
	}
	handler.Events.Append(newEvents...)
	return nil
}

// EventStore abstracts the event store.
type EventStore interface {
	Append(event ...es.Event) error
	Fetch(filter es.EventFilter) ([]es.Event, error)
}

var _ EventStore = &es.Events{}

type TaskCommand interface {
	AggregateID() string
}
