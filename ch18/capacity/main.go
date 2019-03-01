package main

import "fmt"

func main() {
	sattelites := []string{}

	for i := 0; i < 100; i++ {
		sattelites = append(sattelites, "Item")
		fmt.Println(i, cap(sattelites))
	}
}
