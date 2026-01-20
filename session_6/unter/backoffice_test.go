package main

import (
	"math"
	"os"
	"testing"

	"github.com/goccy/go-yaml"
	"github.com/stretchr/testify/require"
)

type priceTest struct {
	Name     string  `yaml:"name"`
	Distance float64 `yaml:"distance"`
	Shared   bool    `yaml:"shared"`
	Want     int     `yaml:"want"`
}

func loadPriceCases(t *testing.T) []priceTest {
	data, err := os.ReadFile("testdata/price_tests.yml")
	require.NoError(t, err, "failed to read test file")

	var priceTests []priceTest
	err = yaml.Unmarshal(data, &priceTests)
	require.NoError(t, err, "failed to parse test file")

	return priceTests
}

func TestRidePrice(t *testing.T) {
	for _, tt := range loadPriceCases(t) {
		t.Run(tt.Name, func(t *testing.T) {
			got := RidePrice(tt.Distance, tt.Shared)
			require.Equal(t, tt.Want, got, "RidePrice(%v, %v)", tt.Distance, tt.Shared)
		})
	}
}

var isCI = os.Getenv("CI") != ""

// In Jenkins use BUILD_NUMBER

func TestInCI(t *testing.T) {
	if !isCI {
		t.Skip("not in CI")
	}

	t.Log("testing")
}

/* Option to run subset of tests
go test -run Example
go test -short # Use testing.Short() + t.Skip in your code
Environment variable
Build tags (//go:build web)
*/

func FuzzRidePrice(f *testing.F) {
	/*
		file, err := os.Create("/tmp/fuzz.txt")
		require.NoError(f, err)
		defer file.Close()
	*/

	f.Add(0.0, true)
	f.Fuzz(func(t *testing.T, distance float64, shared bool) {

		// Sometimes you'll need to adjust the fuzzed data
		distance = math.Abs(distance)

		//fmt.Fprintf(file, "%v %v\n", distance, shared)
		price := RidePrice(distance, shared)
		require.Greater(t, price, 0)
	})
}

// go test -fuzz . -fuzztime 30s
