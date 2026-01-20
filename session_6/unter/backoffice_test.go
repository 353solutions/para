package main

import (
	"fmt"
	"testing"
)

var priceTests = []struct {
	name     string
	distance float64
	shared   bool
	want     int
}{
	{
		name:     "zero distance non-shared",
		distance: 0.0,
		shared:   false,
		want:     250, // initial fare only
	},
	{
		name:     "zero distance shared",
		distance: 0.0,
		shared:   true,
		want:     225, // 250 * 0.9
	},
	{
		name:     "whole number distance non-shared",
		distance: 5.0,
		shared:   false,
		want:     1000, // 250 + (5 * 150)
	},
	{
		name:     "whole number distance shared",
		distance: 5.0,
		shared:   true,
		want:     900, // 1000 * 0.9
	},
	{
		name:     "fractional distance non-shared rounds up",
		distance: 2.3,
		shared:   false,
		want:     700, // 250 + (ceil(2.3)=3 * 150)
	},
	{
		name:     "fractional distance shared rounds up",
		distance: 2.3,
		shared:   true,
		want:     630, // 700 * 0.9
	},
	{
		name:     "single mile non-shared",
		distance: 1.0,
		shared:   false,
		want:     400, // 250 + (1 * 150)
	},
	{
		name:     "single mile shared",
		distance: 1.0,
		shared:   true,
		want:     360, // 400 * 0.9
	},
	{
		name:     "small fractional distance non-shared",
		distance: 0.1,
		shared:   false,
		want:     400, // 250 + (ceil(0.1)=1 * 150)
	},
	{
		name:     "large distance non-shared",
		distance: 20.0,
		shared:   false,
		want:     3250, // 250 + (20 * 150)
	},
	{
		name:     "large distance shared",
		distance: 20.0,
		shared:   true,
		want:     2925, // 3250 * 0.9
	},
}

func TestRidePrice(t *testing.T) {
	for _, tt := range priceTests {
		t.Run(tt.name, func(t *testing.T) {
			got := RidePrice(tt.distance, tt.shared)
			if got != tt.want {
				t.Errorf("RidePrice(%v, %v) = %v, want %v", tt.distance, tt.shared, got, tt.want)
			}
		})
	}

	t.Errorf("oops")
	fmt.Println("Hi")
}
