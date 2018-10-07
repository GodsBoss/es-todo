package main

import (
	"github.com/GodsBoss/es-todo/todo"
)

func task(argsToCommand func([]string) (todo.TaskCommand, error)) func([]string) error {
	return func(args []string) error {
		cmd, err := argsToCommand(args)
		if err != nil {
			return err
		}
		events, err := load()
		if err != nil {
			return err
		}
		(&todo.CommandHandler{
			Events: events,
		}).ProcessTaskCommand(cmd)
		return save(events)
	}
}
