package main

import "fmt"

type celsius float64
type kelvin float64

func celsiusToKelvin(c celsius) kelvin {
	return kelvin(c + 273.15)
}

func main() {
	var c celsius = 127
	k := celsiusToKelvin(c)
	fmt.Println(c, "°C is", k, "°K")
}
