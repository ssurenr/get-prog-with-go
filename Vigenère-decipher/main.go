package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Vigenère Decipher")
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	message := ""

	for i := 0; i < len(cipherText); i++ {
		// A=0, B=1, ... Z=25
		c := cipherText[i] - 'A'
		k := keyword[i%len(keyword)] - 'A'

		// cipher letter - key letter
		c = (c-k+26)%26 + 'A'
		message += string(c)

	}

	fmt.Println(message)
}
