package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	question := "¿Cómo estás?" //Spanish
	fmt.Println(len(question), "bytes")
	fmt.Println(utf8.RuneCountInString(question), "charachters")

	c, size := utf8.DecodeRuneInString(question)
	fmt.Printf("First rune: %c %v bytes", c, size)
}
