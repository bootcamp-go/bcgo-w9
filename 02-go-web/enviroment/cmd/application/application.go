package application

import (
	"app/internal/climate/report"
	"app/internal/climate/storage"
	"errors"
	"fmt"
)

// NewApplication creates a new Application.
func NewApplication(username, storageClimatePath string) *Application {
	return &Application{
		Username:           username,
		StorageClimatePath: storageClimatePath,
	}
}

var (
	// ErrApplicationInternal is returned when an internal error occurs.
	ErrApplicationInternal = errors.New("internal error")
)

// Application represents an application to report climate data.
type Application struct {
	// Username is the one using the application.
	Username string
	// StorageClimatePath is the path to the climate storage.
	StorageClimatePath string
}

func (a *Application) Run() (err error) {
	// dependencies
	stClimate := storage.NewStorageClimateJSON(a.StorageClimatePath)
	rpClimate := report.NewReportClimate(stClimate)

	// run
	// -> get average temperature of countries's climates
	climates, err := rpClimate.AverageTemperature()
	if err != nil {
		err = fmt.Errorf("%w. application failed: %v", ErrApplicationInternal, err)
		return
	}

	// -> report
	fmt.Printf("Welcome %s to the climate report application!\n", a.Username)
	fmt.Printf("The average temperature of the countries is: %.2f\n", climates)

	return
}