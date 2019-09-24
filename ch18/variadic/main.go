package main

import "fmt"

func main() {
	twoWorlds := terraform("New", "Venus", "Mercury")
	fmt.Println(twoWorlds)

	planets := []string{"Venus", "Mars", "Jupiter"}
	newPlanets := terraform("New", planets...)
	fmt.Println()
}

func terraform(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds))

	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}

	return newWorlds
}
