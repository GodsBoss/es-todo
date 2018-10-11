package main

import (
	"github.com/GodsBoss/es-todo/es"
	"github.com/GodsBoss/es-todo/todo"

	"fmt"
	"strconv"
	"time"
)

func toAddCommand(args []string) (es.Command, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("missing task")
	}
	return todo.AddTaskCommand{ID: strconv.Itoa(int(time.Now().Unix())), Description: args[0]}, nil
}
