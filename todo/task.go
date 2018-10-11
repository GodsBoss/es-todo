package todo

import (
	"fmt"

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

func (task *Task) process(command interface{}) ([]es.Event, error) {
	switch cmd := command.(type) {
	case AddTaskCommand:
		return task.add(cmd)
	case RewordTaskCommand:
		return task.reword(cmd)
	case CancelTaskCommand:
		return task.cancel(cmd)
	case FinishTaskCommand:
		return task.finish(cmd)
	case RemoveOldTaskCommand:
		return task.remove(cmd)
	}
	return make([]es.Event, 0), fmt.Errorf("unknown command %+v", command)
}

func (task *Task) add(cmd AddTaskCommand) ([]es.Event, error) {
	if task.state == taskStateInitial {
		return []es.Event{
			es.NewEvent(eventTaskAdded, time.Now(), cmd.ID, []byte(cmd.Description)),
		}, nil
	}
	return noEvents, fmt.Errorf("task already exists")
}

func (task *Task) reword(cmd RewordTaskCommand) ([]es.Event, error) {
	if task.state == taskStateOpen {
		return []es.Event{
			es.NewEvent(eventTaskReworded, time.Now(), cmd.ID, []byte(cmd.Description)),
		}, nil
	}
	return noEvents, fmt.Errorf("task must be open to be reworded")
}

func (task *Task) cancel(cmd CancelTaskCommand) ([]es.Event, error) {
	if task.state == taskStateOpen {
		return []es.Event{
			es.NewEvent(eventTaskCanceled, time.Now(), cmd.ID, []byte{}),
		}, nil
	}
	return noEvents, fmt.Errorf("task must be open to be canceled")
}

func (task *Task) finish(cmd FinishTaskCommand) ([]es.Event, error) {
	if task.state == taskStateOpen {
		return []es.Event{
			es.NewEvent(eventTaskFinished, time.Now(), cmd.ID, []byte{}),
		}, nil
	}
	return noEvents, fmt.Errorf("task must be open to be finished")
}

func (task *Task) remove(cmd RemoveOldTaskCommand) ([]es.Event, error) {
	if task.state == taskStateCanceled || task.state == taskStateFinished {
		return []es.Event{
			es.NewEvent(eventTaskRemoved, time.Now(), cmd.ID, []byte{}),
		}, nil
	}
	return noEvents, fmt.Errorf("task must be closed to be removed")
}

var noEvents = make([]es.Event, 0)

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
