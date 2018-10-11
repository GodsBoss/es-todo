package todo

import "github.com/GodsBoss/es-todo/es"

type CommandHandler struct {
	Events EventStore
}

func (handler *CommandHandler) ProcessTaskCommand(command TaskCommand) error {
	task := &Task{}
	for _, event := range handler.Events.Fetch(es.ByAggregateID(command.AggregateID())) {
		task.apply(event)
	}
	newEvents, err := task.process(command)
	if err != nil {
		return err
	}
	for i := range newEvents {
		handler.Events.Append(newEvents[i])
	}
	return nil
}

// EventStore abstracts the event store.
type EventStore interface {
	Append(event es.Event)
	Fetch(filter es.EventFilter) []es.Event
}

var _ EventStore = &es.Events{}

type TaskCommand interface {
	AggregateID() string
}
