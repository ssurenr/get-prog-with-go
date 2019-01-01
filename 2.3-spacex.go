package main

import "fmt"

func main() {
	const lightSpeed = 299792  // km/s
	const spacexSpeed = 100800 // km/h
	const hoursPerDay = 24
	var distance = 56000000      // km
	var y2025Distance = 96300000 // km

	fmt.Println(distance/lightSpeed, "seconds")

	distance = 401000000

	fmt.Println(y2025Distance/spacexSpeed, "hours")
	fmt.Println(y2025Distance/spacexSpeed/hoursPerDay, "days")
}
