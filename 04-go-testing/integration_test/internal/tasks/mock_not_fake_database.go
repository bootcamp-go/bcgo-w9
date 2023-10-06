package tasks

type MockDatabase struct {
	DataOnGet Task
	ErrOnGet  error
}

func NewMockDatabase(dataOnGet Task, errOnGet error) *MockDatabase {
	return &MockDatabase{dataOnGet, errOnGet}
}

func (m *MockDatabase) Get(id string) (*Task, error) {
	return &m.DataOnGet, m.ErrOnGet
}
