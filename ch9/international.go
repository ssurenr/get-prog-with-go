package main

import "fmt"

func main() {
	message := "Hola Estación Espacial Internacional"

	for _, c := range message {
		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
			c = c + 13
			if c > 'z' || c > 'Z' && c < 'a' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}
	fmt.Println("")
}
