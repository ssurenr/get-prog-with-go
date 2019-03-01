package main

import "fmt"

func main() {
	planets := []string{
		"Mercury", "venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}

	terrestrial := planets[0:4:4]
	worlds := append(terrestrial, "Ceres")
	fmt.Println(planets, len(terrestrial), cap(terrestrial))
	fmt.Println(worlds, len(worlds), cap(worlds))
}
