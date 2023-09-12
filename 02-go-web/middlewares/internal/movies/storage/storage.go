package storage

import "errors"

// Movie is a struct that represents a movie
type Movie struct {
	Id    int
	Title string
	Year  int
	Genre string
}

// StorageMovie is an interface that represents a movie storage
type StorageMovie interface {
	// GetMovies returns all movies
	Get() (m []*Movie, err error)

	// GetByTitle returns a movie by title
	GetByTitle(title string) (m *Movie, err error)
}

var (
	// ErrStorageMovieInternal is returned when an internal error occurs
	ErrStorageMovieInternal = errors.New("internal error")

	// ErrStorageMovieNotFound is returned when a movie is not found
	ErrStorageMovieNotFound = errors.New("movie not found")
)