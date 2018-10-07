package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/GodsBoss/es-todo/todo"
)

func toAddCommand(args []string) (todo.TaskCommand, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("missing task")
	}
	return todo.AddTaskCommand{ID: strconv.Itoa(int(time.Now().Unix())), Description: args[0]}, nil
}
