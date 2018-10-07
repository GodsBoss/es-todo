package main

import (
	"github.com/GodsBoss/es-todo/es/json"
	"github.com/GodsBoss/es-todo/todo"

	"os"
)

func list(_ []string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	events, err := json.FromJSON(file)
	if err != nil {
		return err
	}
	return todo.ListTasks(events).Show(os.Stdout)
}
