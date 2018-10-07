package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/GodsBoss/es-todo/todo"
)

func add(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing task")
	}
	events, err := load()
	if err != nil {
		return err
	}
	(&todo.CommandHandler{
		Events: events,
	}).ProcessTaskCommand(todo.AddTaskCommand{ID: strconv.Itoa(int(time.Now().Unix())), Description: args[0]})
	return save(events)
}
