package main

import (
	"github.com/GodsBoss/es-todo/todo"

	"fmt"
)

func toStateCommand(idToCommand func(id string) todo.TaskCommand) func([]string) (todo.TaskCommand, error) {
	return func(args []string) (todo.TaskCommand, error) {
		if len(args) == 0 {
			return nil, fmt.Errorf("missing ID argument")
		}
		return idToCommand(args[0]), nil
	}
}

func cancel(id string) todo.TaskCommand {
	return todo.CancelTaskCommand{
		ID: id,
	}
}

func finish(id string) todo.TaskCommand {
	return todo.FinishTaskCommand{
		ID: id,
	}
}

func removeOld(id string) todo.TaskCommand {
	return todo.RemoveOldTaskCommand{
		ID: id,
	}
}
