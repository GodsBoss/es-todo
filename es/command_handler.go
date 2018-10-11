package es

import (
	"sync"
)

type CommandHandler interface {
	ProcessCommand(command Command) error
}

type commandHandler struct {
	Events EventStore
	mtx    sync.Mutex
}

func NewCommandHandler(events EventStore) CommandHandler {
	return &commandHandler{
		Events: events,
	}
}

func (handler *commandHandler) ProcessCommand(command Command) error {
	handler.mtx.Lock()
	defer handler.mtx.Unlock()
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
