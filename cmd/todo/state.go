package main

import (
	"github.com/GodsBoss/es-todo/todo"

	"fmt"
)

func toStateCommand(idToCommand func(id string) todo.Command) func([]string) (todo.Command, error) {
	return func(args []string) (todo.Command, error) {
		if len(args) == 0 {
			return nil, fmt.Errorf("missing ID argument")
		}
		return idToCommand(args[0]), nil
	}
}

func cancel(id string) todo.Command {
	return todo.CancelTaskCommand{
		ID: id,
	}
}

func finish(id string) todo.Command {
	return todo.FinishTaskCommand{
		ID: id,
	}
}

func removeOld(id string) todo.Command {
	return todo.RemoveOldTaskCommand{
		ID: id,
	}
}
