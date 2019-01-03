package main

import "fmt"

func main() {
	var room = "cave"

	if room == "cave" {
		fmt.Println("You are in cave")
	} else if room == "entrance" {
		fmt.Println("You are in entrance. Go in")
	} else if room == "mountain" {
		fmt.Println("You are in mountain. Take care")
	} else {
		fmt.Println("Everything is white")
	}
}
