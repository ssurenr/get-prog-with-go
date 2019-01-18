package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var balance uint16 = 0
	var deposit uint16 = 0

	for balance <= 2000 {

		switch rand.Intn(3) {
		case 0:
			deposit = 5
		case 1:
			deposit = 10
		case 2:
			deposit = 25
		}

		balance += deposit
		fmt.Printf("Deposit = $0.%02v, Balance = $%v.%02v\n", deposit, balance/100, balance%100)
	}
}
