package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("You find yourself in a dimly lit cavern.")
	var command = "walk outside"
	var exit = strings.Contains(command, "outside")
	fmt.Println("You leave the cave:", exit)
	// --

	fmt.Println("There is a sign near the entrance that reads 'No Minors'.")

	var age = 41
	var minor = age < 18

	fmt.Printf("At age %v, am I a minor? %v\n", age, minor)
}
