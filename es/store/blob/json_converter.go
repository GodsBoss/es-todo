package blob

import (
	"github.com/GodsBoss/es-todo/es"

	"encoding/json"
	"time"
)

type JSONConverter struct{}

func NewJSONConverter() *JSONConverter {
	return &JSONConverter{}
}

func (converter *JSONConverter) Marshal(events []es.Event) ([]byte, error) {
	jsonEvents := make([]jsonEvent, len(events))
	for i := range events {
		jsonEvents[i] = jsonEvent{
			Timestamp:   events[i].Timestamp(),
			AggregateID: events[i].AggregateID(),
			Kind:        events[i].Kind(),
			Payload:     events[i].Payload(),
		}
	}
	return json.Marshal(jsonEvents)
}

func (converter *JSONConverter) Unmarshal(data []byte) ([]es.Event, error) {
	jsonEvents := make([]jsonEvent, 0)
	err := json.Unmarshal(data, &jsonEvents)
	if err != nil {
		return nil, err
	}
	events := make([]es.Event, len(jsonEvents))
	for i := range jsonEvents {
		events[i] = es.NewEvent(jsonEvents[i].Kind, jsonEvents[i].Timestamp, jsonEvents[i].AggregateID, jsonEvents[i].Payload)
	}
	return events, nil
}

type jsonEvent struct {
	Timestamp   time.Time
	AggregateID string
	Kind        string
	Payload     []byte
}
