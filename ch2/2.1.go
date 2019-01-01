package main

import "fmt"

func main() {
	// fmt.Print("My weight on the surface of Mars is ")
	// fmt.Print(149.0 * 0.3783)
	// fmt.Print(" lbs, and I would be ")
	// fmt.Print(" years old.")
	fmt.Println("My weight on the surface of Mars is", 149.0*0.3783, "lbs, and I would be", 41*365.245/678, "years old.")
	fmt.Printf("My weight on the surface of Mars is %v lbs,", 94*2.2*0.3783)
	fmt.Printf(" and I would be %v years old.\n", 30*365.245/678)

	//Padding
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)
}
