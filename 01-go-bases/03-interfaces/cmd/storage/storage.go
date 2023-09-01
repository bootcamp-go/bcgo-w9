package main

// main
func main() {
	// dependencies
	// -> book
	// -> book storage
	// -> book validator
	stBook := &BookStorageDefault{}
	vlBook := &BookValidatorDefault{}

	// -> book storage validator
	stVlBook := &BookStorageValidator{
		BookStorage: stBook,
		BookValidator: vlBook,
	}

	// -> book
	title := "My Book"
	err := stVlBook.Save(title)
	if err != "" {
		panic(err)
	}
}

// _____________________________________________________________
// BookStorage is the storage interface for books
type BookStorage interface {
	Save(title string) (err string)
}

// _____________________________________________________________
// Implentations
// BookStorageDefault is the book storage						-	#01
type BookStorageDefault struct{}
// Save saves the book
func (s *BookStorageDefault) Save(title string) (err string) {
	return
}

// BookStorageValidator is the book storage validator	-	#02
type BookStorageValidator struct {
	BookStorage
	BookValidator
}
// BookStorageValidator is the book storage validator
func (s *BookStorageValidator) Save(title string) (err string) {
	err = s.BookValidator.Validate(title)
	if err != "" {
		return
	}

	err = s.BookStorage.Save(title)
	return
}

// _____________________________________________________________
// BookValidator is the interface for book validation
type BookValidator interface {
	Validate(title string) (err string)
}

type BookValidatorDefault struct{}
func (v *BookValidatorDefault) Validate(title string) (err string) {
	return
}