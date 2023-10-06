package tasks

type NotFakeDatabase interface {
	Get(id string) (*Task, error)
}

// DefaultNotFakeDatabase is the default implementation of the NotFakeDatabase interface.
type DefaultNotFakeDatabase struct {
	// Data that the database will store.
	Data []Task
}

// Get returns a task by its ID.
// If the task doesn't exist, it returns an empty task.
func (database *DefaultNotFakeDatabase) Get(id string) (result *Task, err error) {
	for index := range database.Data {
		if database.Data[index].ID == id {
			result = &database.Data[index]
			return
		}
	}
	return
}
