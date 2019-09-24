package main

import "fmt"

func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	fmt.Println("Length: ", len(planets))

	for i := 0; i < len(planets); i++ {
		fmt.Println(i, ": ", planets[i])
	}

	for i, planet := range planets {
		fmt.Println(i, "|", planet)
	}

	planetsMarkII := planets

	planets[2] = "whoops"

	fmt.Println(planets)
	fmt.Println(planetsMarkII)
}
