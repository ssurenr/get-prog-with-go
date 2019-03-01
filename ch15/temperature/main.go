package main

import "fmt"

type temperature float64

func drawTable(convert func(temperature) temperature, column1 string, column2 string) {
	// Header
	fmt.Println("===================")
	fmt.Printf("|     °%v |     °%v |\n",column1, column2)
	fmt.Println("===================")

	// Data
	var sample temperature
	for sample = -40; sample <= 100; sample += 5 {
		fmt.Printf("| %6.1f | %6.1f |\n", sample, convert(sample))
	}

	//Footer
	fmt.Println("===================")
}

func celsiusToFarenheit(celsius temperature) temperature {
	return celsius * 9/5 + 32
}

func farenheitToCelsius(farenheit temperature) temperature {
	return (farenheit - 32)*5/9
}

func main() {
	// Celsius Table
	drawTable(celsiusToFarenheit, "C", "F")
	// Farenheit Table
	drawTable(farenheitToCelsius, "F", "C")
}