package report

import (
	"app/internal/climate/storage"
	"errors"
	"fmt"
)

// Report represents a climate report.
func NewReportClimate(st storage.StorageClimate) *ReportClimate {
	return &ReportClimate{st: st}
}

var (
	// ErrReportClimateInternal is returned when an internal error occurs.
	ErrReportClimateInternal = errors.New("internal error")
	// ErrReportClimateNotFound is returned when the climate is not found.
	ErrReportClimateNotFound = errors.New("climate not found")
)

// ReportClimate represents a climate report service.
type ReportClimate struct {
	st storage.StorageClimate
}

func (r *ReportClimate) AverageTemperature() (t float64, err error) {
	// read all climates
	climates, err := r.st.ReadAll()
	if err != nil {
		err = fmt.Errorf("%w. could not read all climates: %v", ErrReportClimateInternal, err)
		return
	}

	// check if there is any climate
	if len(climates) == 0 {
		err = ErrReportClimateNotFound
		return
	}

	// calculate average temperature
	var sum float64
	for _, c := range climates {
		sum += c.AverageTemperature
	}
	t = sum / float64(len(climates))

	return
}