package es

import (
	"time"
)

// Event is an event that already happened.
type Event struct {
	timestamp   time.Time
	aggregateID string
	kind        string
	payload     []byte
}

// NewEvent creates a new event.
func NewEvent(kind string, timestamp time.Time, aggregateID string, payload []byte) Event {
	return Event{
		timestamp:   timestamp,
		aggregateID: aggregateID,
		kind:        kind,
		payload:     payload,
	}
}

// Timestamp returns the event's timestamp.
func (event Event) Timestamp() time.Time {
	return event.timestamp
}

// AggregateID returns the event's aggregate ID.
func (event Event) AggregateID() string {
	return event.aggregateID
}

// Kind returns the event's kind.
func (event Event) Kind() string {
	return event.kind
}

// Payload is the payload of the event. Do not change the return value!
func (event Event) Payload() []byte {
	return event.payload
}

// EventFilter is a filter for events.
type EventFilter func(event Event) bool

// ByKind filters all events by their kind.
func ByKind(kind string) EventFilter {
	return func(event Event) bool {
		return event.kind == kind
	}
}

// ByAggregateID filters all events with the given aggregate ID.
func ByAggregateID(aggregateID string) EventFilter {
	return func(event Event) bool {
		return event.aggregateID == aggregateID
	}
}

func all(_ Event) bool {
	return true
}

// All returns a filter which returns all events.
func All() EventFilter {
	return all
}
