package main

import "fmt"

// KelvinToCelsius converts °K to °C
func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func main() {
	kelvin := 294.0
	kelvin2 := 233.0
	celsius := kelvinToCelsius(kelvin)
	celsius2 := kelvinToCelsius(kelvin2)
	fmt.Println(kelvin,"° K is ",celsius, "° C")
	fmt.Println(kelvin2,"° K is ",celsius2, "° C")
}