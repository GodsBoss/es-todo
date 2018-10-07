package main

import (
	"github.com/GodsBoss/es-todo/es"
	"github.com/GodsBoss/es-todo/es/json"

	"fmt"
	"os"
)

func initFile(_ []string) error {
	file, err := os.Open(filename)
	if err != nil {
		file, err = os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close()
		emptyEvents := &es.Events{}
		err = json.ToJSON(emptyEvents, file)
		return err
	}
	return fmt.Errorf("already initialized")
}

const filename = "events.json"
