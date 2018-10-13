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
	newEvents, err := command.Execute(eventFetcherGuard(handler.Events.Fetch))
	if err != nil {
		return err
	}
	return handler.Events.Append(newEvents...)
}

type EventFetcher interface {
	Fetch(filter EventFilter) ([]Event, error)
}

type eventFetcherGuard func(EventFilter) ([]Event, error)

func (f eventFetcherGuard) Fetch(filter EventFilter) ([]Event, error) {
	return f(filter)
}

// EventStore abstracts the event store.
type EventStore interface {
	EventFetcher

	Append(event ...Event) error
}

type Command interface {
	Execute(EventFetcher) ([]Event, error)
}
