package author

import "fmt"

// Author represents an author of a book
type Author struct {
	FirstName string
	LastName  string
	Age       int
}

// FullName returns the full name of the author
func (a Author) FullName() (fullname string) {
	fullname = fmt.Sprintf("%s %s", a.FirstName, a.LastName)
	return
}

// SetFirstName sets the first name of the author
func (a *Author) SetLastName(lastName string) {
	(*a).LastName = lastName
}