package main

import "fmt"

func main() {
	var days = 28
	var distance = 56000000

	// Speed = distance/time

	const hoursPerDay = 24

	var speed = distance / (days * hoursPerDay)

	fmt.Println("Speed: ", speed)
}
