package todo

import "github.com/GodsBoss/es-todo/es"

type CommandHandler struct {
	Events Events
}

func (handler *CommandHandler) ProcessTaskCommand(command TaskCommand) {
	task := &Task{}
	for _, event := range handler.Events.Fetch(es.ByAggregateID(command.AggregateID())) {
		task.apply(event)
	}
	for _, event := range task.process(command) {
		handler.Events.Append(event)
	}
}

// Events abstracts the event store.
type Events interface {
	Append(event es.Event)
	Fetch(filter es.EventFilter) []es.Event
}

var _ Events = &es.Events{}

type TaskCommand interface {
	AggregateID() string
}
