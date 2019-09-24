package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Printf("%-16v %5v %-10v %-5v\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Printf("=======================================\n")

	distance := 62100000 //km

	for count := 0; count < 10; count++ {
		airline := ""
		duration := 0
		triptype := ""
		price := 0
		const secondsPerDay = 60 * 60 * 24

		switch rand.Intn(3) {
		case 0:
			airline = "Space Adventures"
		case 1:
			airline = "Space X"
		case 2:
			airline = "Virgin Galactic"
		}

		speed := rand.Intn(30-16+1) + 16 // km/s
		duration = distance / speed / secondsPerDay

		price = 20.0 + speed

		if rand.Intn(2) == 0 {
			triptype = "One-way"
		} else {
			triptype = "Two-way"
			price *= 2
		}

		fmt.Printf("%-16v %5v %-10v $%4v\n", airline, duration, triptype, price)
	}
}
