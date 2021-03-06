package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func main() {
	for count := 0; count < 10; count++ {
		// year := rand.Intn(2018) + 1
		year := 2019 + rand.Intn(10)
		month := rand.Intn(12) + 1
		daysInMonth := 31

		switch month {
		case 2:
			leap := year%400 == 0 || (year%4 == 0 && year%100 != 0)
			// if leap == false {
			// 	daysInMonth = 28
			// } else {
			// 	daysInMonth = 29
			// }
			daysInMonth = 28
			if leap {
				daysInMonth = 29
			}
		case 4, 6, 9, 11:
			daysInMonth = 30
		}

		day := rand.Intn(daysInMonth) + 1
		fmt.Println(era, year, month, day)
	}

}
