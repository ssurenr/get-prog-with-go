package main

import (
	"fmt"
)

func main() {
	planets := [...]string{
		"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune",
	}

	terrestial := planets[0:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:8]

	fmt.Println("Terrestial: ", terrestial)
	fmt.Println("gasGiants: ", gasGiants)
	fmt.Println("iceGiants: ", iceGiants)

	fmt.Println(gasGiants[1])

	giants := planets[4:8]
	gas := giants[0:2]
	ice := giants[2:4]
	fmt.Println(giants, gas, ice)

	iceGiantsMarkII := iceGiants
	iceGiants[1] = "Poseidon"
	fmt.Println(planets)
	fmt.Println(iceGiants, iceGiantsMarkII, ice)

	// default indices

	terrestial = planets[:4]
	gasGiants = planets[4:6]
	iceGiants = planets[6:]

	fmt.Println("Terrestial: ", terrestial)
	fmt.Println("gasGiants: ", gasGiants)
	fmt.Println("iceGiants: ", iceGiants)

	colonized := terrestial[2:]
	fmt.Println(colonized)

	dwarfArray := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfSlice := dwarfArray[:]

	fmt.Printf("DwarfArray: %T, DwarfSlice: %T", dwarfArray, dwarfSlice)
}
