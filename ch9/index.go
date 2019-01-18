package main

import "fmt"

func main() {
	message := "Shalom"
	c := message[5]
	message[5] = 'n'
	fmt.Printf("%c\n", c)
}
