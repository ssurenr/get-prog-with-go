package main

import "fmt"

func main() {
	planets := map[string]string{"Earth": "Sector 779", "Mars": "Sector 779"}

	planetsMarkII := planets

	planets["Earth"] = "whoops"
	fmt.Println(planets)
	fmt.Println(planetsMarkII)
	delete(planets, "Earth")
	fmt.Println(planetsMarkII)
}
