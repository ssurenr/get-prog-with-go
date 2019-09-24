package main

import (
	"fmt"
)

func main() {
	type location struct {
		lat float64
		long float64
	}

	var spirit location
	spirit.lat = -14.5684
	spirit.long = 175.472636

	var oppurtunity location
	oppurtunity.lat = -1.9462
	oppurtunity.long  =354.4734

	fmt.Println(spirit, oppurtunity)
}
