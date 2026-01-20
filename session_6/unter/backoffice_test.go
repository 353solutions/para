package main

import (
	"os"
	"testing"

	"github.com/goccy/go-yaml"
)

type priceTest struct {
	Name     string  `yaml:"name"`
	Distance float64 `yaml:"distance"`
	Shared   bool    `yaml:"shared"`
	Want     int     `yaml:"want"`
}

func TestRidePrice(t *testing.T) {
	data, err := os.ReadFile("testdata/price_tests.yml")
	if err != nil {
		t.Fatalf("failed to read test file: %v", err)
	}

	var priceTests []priceTest
	if err := yaml.Unmarshal(data, &priceTests); err != nil {
		t.Fatalf("failed to parse test file: %v", err)
	}

	for _, tt := range priceTests {
		t.Run(tt.Name, func(t *testing.T) {
			got := RidePrice(tt.Distance, tt.Shared)
			if got != tt.Want {
				t.Errorf("RidePrice(%v, %v) = %v, want %v", tt.Distance, tt.Shared, got, tt.Want)
			}
		})
	}
}
