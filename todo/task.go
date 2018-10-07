package todo

import (
	"github.com/GodsBoss/es-todo/es"

	"time"
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

func (task *Task) process(command interface{}) []es.Event {
	switch cmd := command.(type) {
	case AddTaskCommand:
		if task.state == taskStateInitial {
			return []es.Event{
				es.NewEvent(eventTaskAdded, time.Now(), cmd.ID, []byte(cmd.Description)),
			}
		}
	case RewordTaskCommand:
		if task.state == taskStateOpen {
			return []es.Event{
				es.NewEvent(eventTaskReworded, time.Now(), cmd.ID, []byte(cmd.Description)),
			}
		}
	case CancelTaskCommand:
		if task.state == taskStateOpen {
			return []es.Event{
				es.NewEvent(eventTaskCanceled, time.Now(), cmd.ID, []byte{}),
			}
		}
	case FinishTaskCommand:
		if task.state == taskStateOpen {
			return []es.Event{
				es.NewEvent(eventTaskFinished, time.Now(), cmd.ID, []byte{}),
			}
		}
	case RemoveOldTaskCommand:
		if task.state == taskStateCanceled || task.state == taskStateFinished {
			return []es.Event{
				es.NewEvent(eventTaskRemoved, time.Now(), cmd.ID, []byte{}),
			}
		}
	}
	return make([]es.Event, 0)
}

const (
	taskStateInitial  = ""
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
