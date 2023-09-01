package main

import (
	"app/internal/author"
	"app/internal/book"
	"fmt"
)

func main() {
	// ptrs
	var v int = 1
	c := &v
	fmt.Printf("v: %p\nc: %p\n", &v, c)

	// dependencies
	// -> author
	a := author.Author{
		FirstName: "John",
		LastName:  "Doe",
		Age:       25,
	}
	a.SetLastName("Smith")

	// -> book
	b := book.Book{
		Title:  "My Book",
		Description: "My Description",
		Author: a,
	}
	b.FullInfo()
}
		