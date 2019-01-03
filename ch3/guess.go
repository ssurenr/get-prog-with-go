package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var myNumber = 63

	for {
		var computerGuess = rand.Intn(100) + 1

		if computerGuess > myNumber {
			fmt.Printf("Your number %v is too high. Try again.\n", computerGuess)
		} else if computerGuess < myNumber {
			fmt.Printf("your number %v is too low. Try again.\n", computerGuess)
		} else {
			fmt.Println("You got it right! The number is ", computerGuess)
			break
		}
	}
}
