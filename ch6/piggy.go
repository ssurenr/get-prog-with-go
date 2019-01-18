package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var balance float64 = 0.0
	var deposit float64 = 0.0

	for balance <= 20.0 {

		switch rand.Intn(3) {
		case 0:
			deposit = 0.05
		case 1:
			deposit = 0.10
		case 2:
			deposit = 0.25
		}

		balance += deposit
		fmt.Printf("Deposit = %4.2f, Balance = %5.2f\n", deposit, balance)
	}
}
