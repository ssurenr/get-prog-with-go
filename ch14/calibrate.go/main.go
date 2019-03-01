package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

//sensor function type
type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func fakeSensor() kelvin {
	return kelvin(150 + rand.Intn(151))
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin{
		return s() + offset
	}
}

func main() {
	var errorCorrection kelvin = 5
	sensor := calibrate(fakeSensor, errorCorrection)
	for i := 0; i < 10; i++ {
		fmt.Println(sensor())
	}
}
