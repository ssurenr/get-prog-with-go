package main

import (
	"fmt"
	"sort"
)

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth",
		"Mars", "Jupiter", "Saturn",
		"Uranus", "Neptune",
	}

	//sort.StringSlice(planets).Sort()
	sort.Strings(planets)
	fmt.Printf("planets - %T\n", planets)
	slice := sort.StringSlice(planets)
	fmt.Printf("slice - %T\n", slice)
}
