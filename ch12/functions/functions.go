package main

import "fmt"

// KelvinToCelsius converts °K to °C
func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func celsiusToFarenheit(c float64) float64 {
	f := (c * 9.0 / 5.0) + 32
	return f
}

func kelvinToFarenheit(k float64) float64 {
	c := kelvinToCelsius(k)
	f := celsiusToFarenheit(c)
	return f
}

func main() {
	kelvin := 0.0
	farenheit := kelvinToFarenheit(kelvin)
	fmt.Println(kelvin, "° K is ", farenheit, "° F")
}
