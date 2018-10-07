package main

import (
	"github.com/GodsBoss/es-todo/todo"

	"fmt"
)

func toRewordCommand(args []string) (todo.TaskCommand, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("needs ID and task as args")
	}
	return todo.RewordTaskCommand{
		ID:          args[0],
		Description: args[1],
	}, nil
}
