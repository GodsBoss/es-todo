package main

import (
	"github.com/GodsBoss/es-todo/todo"

	"os"
)

func list(_ []string) error {
	events, err := load()
	if err != nil {
		return err
	}
	return todo.ListTasks(events).Show(os.Stdout)
}
