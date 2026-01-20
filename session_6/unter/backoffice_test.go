package main

import (
	"testing"
)

func TestRidePrice(t *testing.T) {
	// TODO: Load test cases from testdata/price_cases.yml
	for _, tt := range priceTests {

		t.Run(tt.name, func(t *testing.T) {
			got := RidePrice(tt.distance, tt.shared)
			if got != tt.want {
				t.Errorf("RidePrice(%v, %v) = %v, want %v", tt.distance, tt.shared, got, tt.want)
			}
		})
	}
}
