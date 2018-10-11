package main

import (
	"github.com/GodsBoss/es-todo/es/store/blob"
)

const filename = "events.json"

func eventStore() *blob.EventStore {
	return blob.NewEventStore(
		blob.NewFileBlobStore(filename),
		blob.NewJSONConverter(),
	)
}
