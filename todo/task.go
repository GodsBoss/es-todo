package todo

import (
	"github.com/GodsBoss/es-todo/es"
)

type Task struct {
	id          string
	description string
	state       string
}

func (task *Task) apply(event es.Event) {
	switch event.Kind() {
	case eventTaskAdded:
		task.id = event.AggregateID()
		task.state = taskStateOpen
		task.description = string(event.Payload())
	case eventTaskReworded:
		task.description = string(event.Payload())
	case eventTaskCanceled:
		task.state = taskStateCanceled
	case eventTaskFinished:
		task.state = taskStateFinished
	case eventTaskRemoved:
		task.state = taskStateRemoved
	}
}

const (
	taskStateOpen     = "open"
	taskStateCanceled = "canceled"
	taskStateFinished = "finished"
	taskStateRemoved  = "removed"
)

const (
	eventTaskAdded    = "task_added"
	eventTaskReworded = "task_reworded"
	eventTaskCanceled = "task_canceled"
	eventTaskFinished = "task_finished"
	eventTaskRemoved  = "task_removed"
)
