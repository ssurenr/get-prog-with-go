package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var count = 10

	for count > 0 {
		if rand.Intn(100) == 0 {
			break
		}

		fmt.Println(count)
		time.Sleep(time.Second)
		count--
	}

	if count == 0 {
		fmt.Println("Liftoff!")
	} else {
		fmt.Println("Launch failed")
	}
}
