package main

import (
	"github.com/GodsBoss/es-todo/es"
	"github.com/GodsBoss/es-todo/todo"

	"fmt"
)

func toStateCommand(idToCommand func(id string) es.Command) func([]string) (es.Command, error) {
	return func(args []string) (es.Command, error) {
		if len(args) == 0 {
			return nil, fmt.Errorf("missing ID argument")
		}
		return idToCommand(args[0]), nil
	}
}

func cancel(id string) es.Command {
	return todo.CancelTaskCommand{
		ID: id,
	}
}

func finish(id string) es.Command {
	return todo.FinishTaskCommand{
		ID: id,
	}
}

func removeOld(id string) es.Command {
	return todo.RemoveOldTaskCommand{
		ID: id,
	}
}
