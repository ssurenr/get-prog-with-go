package main

import "fmt"

func main() {
	message := "shalom"

	for i := 0; i < 6; i++ {
		fmt.Printf("%c\n", message[i])
	}
}
