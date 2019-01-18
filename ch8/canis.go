package main

import "fmt"

func main() {
	const distance = 236e15   // km
	const lightSpeed = 299792 // km/s
	const secondsPerYear = 60 * 60 * 24 * 365

	const lightYear = distance / lightSpeed / secondsPerYear

	fmt.Println("Canis Major Dwarf is ", lightYear, " light years away.")
}
