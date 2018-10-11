package main

import (
	"github.com/GodsBoss/es-todo/todo"

	"os"
)

func list(_ []string) error {
	return todo.ListTasks(eventStore()).Show(os.Stdout)
}
