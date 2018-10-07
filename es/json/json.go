package json

import (
	"github.com/GodsBoss/es-todo/es"

	"encoding/json"
	"io"
	"time"
)

func ToJSON(source *es.Events, output io.Writer) error {
	events := source.Fetch(es.All())
	jsonEvents := make([]Event, len(events))
	for i := range events {
		jsonEvents[i] = Event{
			Timestamp:   events[i].Timestamp(),
			AggregateID: events[i].AggregateID(),
			Kind:        events[i].Kind(),
			Payload:     events[i].Payload(),
		}
	}
	return json.NewEncoder(output).Encode(jsonEvents)
}

func FromJSON(input io.Reader) (*es.Events, error) {
	jsonEvents := make([]Event, 0)
	err := json.NewDecoder(input).Decode(&jsonEvents)
	if err != nil {
		return nil, err
	}
	events := &es.Events{}
	for i := range jsonEvents {
		events.Append(es.NewEvent(jsonEvents[i].Kind, jsonEvents[i].Timestamp, jsonEvents[i].AggregateID, jsonEvents[i].Payload))
	}
	return events, nil
}

type Event struct {
	Timestamp   time.Time
	AggregateID string
	Kind        string
	Payload     []byte
}
