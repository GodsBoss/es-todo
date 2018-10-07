package todo

// AddTaskCommand adds a new task, identified by ID. It must not exist yet.
type AddTaskCommand struct {
	ID          string
	Description string
}

// RewordTaskCommand rewords a task, identified by ID. The task must exist and still be open.
type RewordTaskCommand struct {
	ID          string
	Description string
}

// CancelTaskCommand cancels a task, identified by ID. The task must exist and still be open.
type CancelTaskCommand struct {
	ID string
}

// FinishTaskCommand finishes a task, identified by ID. The task must exist and still be open.
type FinishTaskCommand struct {
	ID string
}

// RemoveOldTaskCommand removes a task, identified by ID. The task must exist and not be open anymore.
type RemoveOldTaskCommand struct {
	ID string
}
