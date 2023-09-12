package storage

import "fmt"

// NewStorageMovieInMem returns a new instance of StorageMovieInMem
func NewStorageMovieInMem(db []*Movie) *StorageMovieInMem {
	return &StorageMovieInMem{
		db: db,
	}
}

// StorageMovieInMem is a struct that represents a movie storage in memory
type StorageMovieInMem struct {
	db []*Movie
}

// Get returns all movies
func (s *StorageMovieInMem) Get() (m []*Movie, err error) {
	// do a copy of the slice
	m = make([]*Movie, len(s.db))
	copy(m, s.db)

	return
}

// GetByTitle returns a movie by title
func (s *StorageMovieInMem) GetByTitle(title string) (m *Movie, err error) {
	var exists bool
	for i := range s.db {
		if s.db[i].Title == title {
			m = s.db[i]
			exists = true
			break
		}
	}

	if !exists {
		err = fmt.Errorf("%w: %s", ErrStorageMovieNotFound, title)
		return
	}

	return
}
