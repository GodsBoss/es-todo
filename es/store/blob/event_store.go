package blob

import (
	"github.com/GodsBoss/es-todo/es"

	"fmt"
)

type EventStore struct {
	blobStore BlobStore
	converter Converter
	events    []es.Event
}

func NewEventStore(blobStore BlobStore, converter Converter) *EventStore {
	return &EventStore{
		blobStore: blobStore,
		converter: converter,
	}
}

type BlobStore interface {
	// Load returns a store blob previously stored. If data was never stored, the ByteStore must return an error created with NewNotFoundError().
	Load() ([]byte, error)

	// Save stores the events of the EventStore in the byte store.
	Save([]byte) error
}

type Converter interface {
	Marshal([]es.Event) ([]byte, error)
	Unmarshal([]byte) ([]es.Event, error)
}

type notFoundError struct {
	message string
}

func (err *notFoundError) Error() string {
	return err.message
}

func NewNotFoundError(message string) error {
	return &notFoundError{
		message: message,
	}
}

func IsNotFoundError(err error) bool {
	_, ok := err.(*notFoundError)
	return ok
}

func (store *EventStore) Init() error {
	_, err := store.blobStore.Load()
	if err == nil {
		return fmt.Errorf("event store already initialized")
	}
	if !IsNotFoundError(err) {
		return err
	}
	store.events = make([]es.Event, 0)
	return store.save()
}

func (store *EventStore) Append(events ...es.Event) error {
	if err := store.load(); err != nil {
		return err
	}
	store.events = append(store.events, events...)
	return store.save()
}

func (store *EventStore) Fetch(filter es.EventFilter) ([]es.Event, error) {
	if err := store.load(); err != nil {
		return nil, err
	}
	events := make([]es.Event, 0)
	for i := range store.events {
		events = append(events, store.events[i])
	}
	return events, nil
}

// save stores all events of the event store in the byte store.
func (store *EventStore) save() error {
	data, err := store.converter.Marshal(store.events)
	if err != nil {
		return err
	}
	return store.blobStore.Save(data)
}

// load loads all events found in the byte store, but only if not loaded yet.
func (store *EventStore) load() error {
	if store.events != nil {
		return nil
	}
	data, err := store.blobStore.Load()
	if err != nil {
		return err
	}
	events, err := store.converter.Unmarshal(data)
	if err != nil {
		return err
	}
	store.events = events
	return nil
}
