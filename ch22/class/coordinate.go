package main

import (
	"fmt"
)

type coordinate struct {
	degree, minute, second float64
	h                      rune
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.degree + c.minute/60 + c.second/3600)
}

func main() {
	// Bradbury landing: 4º35'22.2" S, 137º26'30.1" E
	bradLat := coordinate{4, 33, 22.2, 'S'}
	bradLong := coordinate{137, 26, 30.1, 'E'}

	fmt.Printf("Bradbury Decimal Coordinates: %.4f, %.4f\n", bradLat.decimal(), bradLong.decimal())

	// Bradbury variable from constructor
	bradbury := newLocation(coordinate{4, 35, 22.2, 'S'},
		coordinate{137, 26, 30.12, 'E'})
	fmt.Printf("Bradbury Decimal Coordinates: %.4f, %.4f\n", bradbury.lat, bradbury.long)

	var mars = world{radius: 3389.5}
	spirit := location{-14.5684, 175.472636}
	opportunity := location{
		lat:  -1.9462,
		long: 354.4734,
	}

	distance := mars.distance(spirit, opportunity)
	fmt.Printf("Spirit - Opportunity Seperation = %.2f km\n", distance)

}
