package todo

import (
	"github.com/GodsBoss/es-todo/es"

	"fmt"
	"io"
)

type ListTask struct {
	State       string
	Description string
}

func (lt ListTask) String() string {
	stateString := " "
	if lt.State == taskStateFinished {
		stateString = "✓"
	}
	if lt.State == taskStateCanceled {
		stateString = "✗"
	}
	return fmt.Sprintf("[%s] %s", stateString, lt.Description)
}

type Tasks map[string]*ListTask

func (tasks Tasks) Show(output io.Writer) error {
	for id, task := range tasks {
		_, err := fmt.Fprintf(output, "%10s: %s\n", id, task)
		if err != nil {
			return err
		}
	}
	return nil
}

func ListTasks(events es.EventFetcher) Tasks {
	tasks := make(Tasks)
	evs, _ := events.Fetch(es.All())
	for i := range evs {
		event := evs[i]
		switch event.Kind() {
		case eventTaskAdded:
			tasks[event.AggregateID()] = &ListTask{
				State:       taskStateOpen,
				Description: string(event.Payload()),
			}
		case eventTaskReworded:
			tasks[event.AggregateID()].Description = string(event.Payload())
		case eventTaskCanceled:
			tasks[event.AggregateID()].State = taskStateCanceled
		case eventTaskFinished:
			tasks[event.AggregateID()].State = taskStateFinished
		case eventTaskRemoved:
			delete(tasks, event.AggregateID())
		}
	}
	return tasks
}
