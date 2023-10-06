package tasks

import "errors"

var (
	// ErrTaskNotFound is an error that indicates that a task doesn't exist.
	ErrTaskNotFound = errors.New("task not found")
)

// Repository is a storage layer for tasks.
type Repository interface {
	// GetByID returns a task by its ID.
	// If the task doesn't exist, it returns an error.
	GetByID(id string) (Task, error)
}

// DefaultRepository is the default implementation of the Repository interface.
type DefaultRepository struct {
	// Database is the storage layer for tasks.
	Database NotFakeDatabase
}

// GetByID returns a task by its ID.
// If the task doesn't exist, it returns an error.
func (repository *DefaultRepository) GetByID(id string) (result Task, err error) {
	// Try to get the taskFound from the database.
	taskFound, err := repository.Database.Get(id)
	if err != nil {
		return
	}

	// If the task doesn't exist, return an error.
	if taskFound == nil {
		err = ErrTaskNotFound
		return
	}

	// Return the task.
	result = *taskFound
	return
}
