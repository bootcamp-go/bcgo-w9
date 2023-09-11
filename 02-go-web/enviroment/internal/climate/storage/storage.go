package storage

import "errors"

// Climate represents a climate data.
type Climate struct {
	Country 		   string
	AverageTemperature float64
}

type StorageClimate interface {
	// ReadAll reads all the data from the storage.
	ReadAll() (c []*Climate, err error)
}

var (
	// ErrStorageClimateInternal is returned when an internal error occurs.
	ErrStorageClimateInternal = errors.New("internal error")
)