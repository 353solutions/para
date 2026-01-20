package main

import (
	"fmt"
	"time"
)

type Ride struct {
	ID         string
	Time       time.Time
	DistanceKM float64
	Shared     bool
}

func (r Ride) Validate() error {
	if r.ID == "" {
		return fmt.Errorf("empty ID")
	}

	if r.Time.IsZero() {
		return fmt.Errorf("missing time")
	}

	if r.DistanceKM <= 0 {
		return fmt.Errorf("%f - bad distance", r.DistanceKM)
	}

	return nil
}
