package book

import (
	"app/internal/author"
	"fmt"
)

// Book represents a book
type Book struct {
	Title  		string
	Description string
	Author 		author.Author
}

func (b Book) FullInfo() string {
	return fmt.Sprintf("%s - %s - %s", b.Title, b.Description, b.Author.FullName())
}