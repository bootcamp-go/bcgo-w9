package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

// NewStorageClimateJSON creates a new StorageClimateJSON.
func NewStorageClimateJSON(path string) *StorageClimateJSON {
	return &StorageClimateJSON{path: path}
}

// StorageClimateJSON represents a storage for climate data in JSON format.
type StorageClimateJSON struct {
	path string
}

// ClimateJSON represents a struct for climate data in JSON format.
type ClimateJSON struct {
	Country            string  `json:"country"`
	AverageTemperature float64 `json:"average_temperature"`
}

// ReadAll reads all the data from the storage.
func (s *StorageClimateJSON) ReadAll() (c []*Climate, err error) {
	// open file
	f, err := os.Open(s.path)
	if err != nil {
		err = fmt.Errorf("%w. could not open the file: %v", ErrStorageClimateInternal, err)
		return
	}
	defer f.Close()

	// read file
	var climatesJSON []*ClimateJSON
	err = json.NewDecoder(f).Decode(&climatesJSON)
	if err != nil {
		err = fmt.Errorf("%w. could not decode the file: %v", ErrStorageClimateInternal, err)
		return
	}

	// serialization
	for _, cj := range climatesJSON {
		climate := &Climate{
			Country:            cj.Country,
			AverageTemperature: cj.AverageTemperature,
		}

		c = append(c, climate)
	}
	
	return
}