package main

import "fmt"

type planets []string

func (p planets) terraform() {
	for i, planet := range p{
		p[i] = "New " + planet
	}
}

func main() {
	var randomPlanets planets= []string{"Mars", "Uranus", "Neptune"}
	fmt.Println(randomPlanets)
	randomPlanets.terraform()
	fmt.Println(randomPlanets)
}
