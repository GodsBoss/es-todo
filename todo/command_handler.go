package todo

import "github.com/GodsBoss/es-todo/es"

type CommandHandler struct {
	Events Events
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

// Events abstracts the event store.
type Events interface {
	Append(event es.Event)
	Fetch(filter es.EventFilter) []es.Event
}

var _ Events = &es.Events{}

type TaskCommand interface {
	AggregateID() string
}
