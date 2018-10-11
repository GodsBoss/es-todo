package todo

import (
	"github.com/GodsBoss/es-todo/es"
)

// AddTaskCommand adds a new task, identified by ID. It must not exist yet.
type AddTaskCommand struct {
	ID          string
	Description string
}

func (cmd AddTaskCommand) Execute(fetcher EventFetcher) ([]es.Event, error) {
	task, err := loadTask(fetcher, cmd.ID)
	if err != nil {
		return nil, err
	}
	return task.add(cmd)
}

// RewordTaskCommand rewords a task, identified by ID. The task must exist and still be open.
type RewordTaskCommand struct {
	ID          string
	Description string
}

func (cmd RewordTaskCommand) Execute(fetcher EventFetcher) ([]es.Event, error) {
	task, err := loadTask(fetcher, cmd.ID)
	if err != nil {
		return nil, err
	}
	return task.reword(cmd)
}

// CancelTaskCommand cancels a task, identified by ID. The task must exist and still be open.
type CancelTaskCommand struct {
	ID string
}

func (cmd CancelTaskCommand) Execute(fetcher EventFetcher) ([]es.Event, error) {
	task, err := loadTask(fetcher, cmd.ID)
	if err != nil {
		return nil, err
	}
	return task.cancel(cmd)
}

// FinishTaskCommand finishes a task, identified by ID. The task must exist and still be open.
type FinishTaskCommand struct {
	ID string
}

func (cmd FinishTaskCommand) Execute(fetcher EventFetcher) ([]es.Event, error) {
	task, err := loadTask(fetcher, cmd.ID)
	if err != nil {
		return nil, err
	}
	return task.finish(cmd)
}

// RemoveOldTaskCommand removes a task, identified by ID. The task must exist and not be open anymore.
type RemoveOldTaskCommand struct {
	ID string
}

func (cmd RemoveOldTaskCommand) Execute(fetcher EventFetcher) ([]es.Event, error) {
	task, err := loadTask(fetcher, cmd.ID)
	if err != nil {
		return nil, err
	}
	return task.remove(cmd)
}

func loadTask(fetcher EventFetcher, id string) (*Task, error) {
	pastEvents, err := fetcher.Fetch(es.ByAggregateID(id))
	if err != nil {
		return nil, err
	}
	task := &Task{}
	for i := range pastEvents {
		task.apply(pastEvents[i])
	}
	return task, nil
}
