package main

import "fmt"

func main() {
	quote := "L fdph, L vdz, L frqtxhuhg."

	for i := 0; i < len(quote); i++ {
		char := quote[i]
		if char >= 'A' && char <= 'Z' || char >= 'a' && char <= 'z' {
			char = char - 3
			if char < 'A' || char < 'a' && char > 'Z' {
				char += 26
			}
		}
		fmt.Printf("%c", char)
	}
	fmt.Println("")
}
