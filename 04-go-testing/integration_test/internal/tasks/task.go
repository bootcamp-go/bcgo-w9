package tasks

// Task is a structure that represents a task in the system,
// containing all the information needed to execute it.
type Task struct {
	// ID is a unique identifier for the task, in UUID v4 format.
	ID string `json:"id"`

	// Name of the task.
	Name string `json:"name"`

	// Description of the task.
	Description string `json:"description"`

	// Completed is a boolean that indicates if the task is completed or not.
	Completed bool `json:"completed"`
}
