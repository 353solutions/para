package main

import "math"

const sharedDiscount = 0.1

// RidePrice returns ride price in ¢
func RidePrice(distance float64, shared bool) int {
	price := 250 // initial fare
	price += int(math.Ceil(distance)) * 150

	if shared {
		price = int(float64(price) * (1 - sharedDiscount))
	}

	return price
}
