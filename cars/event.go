package main

import (
	"fmt"
	"time"
)

type Event struct {
	ID   string    `json:"id"`
	Time time.Time `json:"start"`
	Lat  float64   `json:"lat"`
	Lng  float64   `json:"lng"`
}

func (r Event) Validate() error {
	if r.ID == "" {
		return fmt.Errorf("empty ID")
	}

	if r.Time.IsZero() {
		return fmt.Errorf("missing Start")
	}

	if r.Lat < -90 || r.Lat > 90 {
		return fmt.Errorf("bad lat")
	}

	if r.Lng < -180 || r.Lng > 180 {
		return fmt.Errorf("bad lng")
	}

	return nil
}
