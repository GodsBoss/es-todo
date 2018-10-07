package main

import (
	"github.com/GodsBoss/es-todo/todo"

	"fmt"
)

func reword(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("needs ID and task as args")
	}
	events, err := load()
	if err != nil {
		return err
	}
	(&todo.CommandHandler{
		Events: events,
	}).ProcessTaskCommand(todo.RewordTaskCommand{
		ID:          args[0],
		Description: args[1],
	})
	return save(events)
}
