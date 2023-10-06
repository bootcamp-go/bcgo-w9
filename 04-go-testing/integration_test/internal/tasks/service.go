package tasks

// Service is a bussiness logic layer for managing tasks in the system.
type Service interface {
	// GetByID returns a task by its ID.
	// If the task doesn't exist, it returns an error.
	GetByID(id string) (Task, error)
}

// DefaultService is the default implementation of the Service interface.
type DefaultService struct {
	// Repository is the storage layer for tasks.
	Repository Repository
}

// GetByID returns a task by its ID.
// If the task doesn't exist, it returns an error.
func (service *DefaultService) GetByID(id string) (result Task, err error) {
	// Try to get the task from the repository.
	taskFound, err := service.Repository.GetByID(id)
	if err != nil {
		return
	}

	// Return the task.
	result = taskFound
	return
}
