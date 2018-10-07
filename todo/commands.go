package todo

// AddTaskCommand adds a new task, identified by ID. It must not exist yet.
type AddTaskCommand struct {
	ID          string
	Description string
}

func (cmd AddTaskCommand) AggregateID() string {
	return cmd.ID
}

// RewordTaskCommand rewords a task, identified by ID. The task must exist and still be open.
type RewordTaskCommand struct {
	ID          string
	Description string
}

func (cmd RewordTaskCommand) AggregateID() string {
	return cmd.ID
}

// CancelTaskCommand cancels a task, identified by ID. The task must exist and still be open.
type CancelTaskCommand struct {
	ID string
}

func (cmd CancelTaskCommand) AggregateID() string {
	return cmd.ID
}

// FinishTaskCommand finishes a task, identified by ID. The task must exist and still be open.
type FinishTaskCommand struct {
	ID string
}

func (cmd FinishTaskCommand) AggregateID() string {
	return cmd.ID
}

// RemoveOldTaskCommand removes a task, identified by ID. The task must exist and not be open anymore.
type RemoveOldTaskCommand struct {
	ID string
}

func (cmd RemoveOldTaskCommand) AggregateID() string {
	return cmd.ID
}
