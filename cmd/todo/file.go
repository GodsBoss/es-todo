package main

import (
	"github.com/GodsBoss/es-todo/es"
	"github.com/GodsBoss/es-todo/es/json"

	"os"
)

const filename = "events.json"

func load() (*es.Events, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return json.FromJSON(file)
}
