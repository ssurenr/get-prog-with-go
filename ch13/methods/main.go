package main

import "fmt"

type celsius float64
type farenheit float64
type kelvin float64

func (c celsius) farenheit() farenheit {
	return farenheit(c * 9/5 + 32)
}

func (c celsius) kelvin() kelvin {
	return kelvin(273.15 + c)
}

func (f farenheit) celsius() celsius {
	return celsius((f - 32)/1.8)
}

func (f farenheit) kelvin() kelvin {
	return kelvin((f + 459.67)*5/9)
}

func (k kelvin) celsius() celsius  {
	return celsius(k - 273.15)
}

func (k kelvin) farenheit() farenheit {
	return farenheit((k + 459.67) * 5/9)
}

func main() {
	var temp celsius = 34
	fmt.Println(temp, "°C = ", temp.farenheit(), "°F and ", temp.kelvin(), "K")
}
