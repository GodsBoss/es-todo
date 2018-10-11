package main

import (
	"github.com/GodsBoss/es-todo/es"
)

func task(argsToCommand func([]string) (es.Command, error)) func([]string) error {
	return func(args []string) error {
		cmd, err := argsToCommand(args)
		if err != nil {
			return err
		}
		events, err := load()
		if err != nil {
			return err
		}
		err = (&es.CommandHandler{
			Events: events,
		}).ProcessCommand(cmd)
		if err != nil {
			return err
		}
		return save(events)
	}
}
