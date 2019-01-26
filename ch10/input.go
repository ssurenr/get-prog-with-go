package main

import "fmt"

func main() {
	var input string = "1"
	var decision bool

	switch input {
	case "true", "yes", "1":
		decision = true
		fmt.Println(decision)
	case "false", "no", "0":
		decision = false
		fmt.Println(decision)

	default:
		fmt.Println("Acceptable values true/false, yes/no, 1/0")
		return
	}

}
