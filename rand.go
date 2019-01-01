package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num = rand.Intn(10) + 1
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)

	//Generate a random number between 56,000,000 to 401,000,000
	// 👇🏻  +1 is for fixing off-by-one error
	num = rand.Intn(401000000-56000000+1) + 56000000
	fmt.Println(num)
	fmt.Println("401000000-56000000=", 401000000-56000000)
}
