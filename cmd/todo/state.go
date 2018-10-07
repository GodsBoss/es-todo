package main

import (
	"github.com/GodsBoss/es-todo/todo"

	"fmt"
)

func state(idToCommand func(id string) todo.TaskCommand) func([]string) error {
	return func(args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("missing ID argument")
		}
		events, err := load()
		if err != nil {
			return err
		}
		(&todo.CommandHandler{
			Events: events,
		}).ProcessTaskCommand(idToCommand(args[0]))
		return save(events)
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
