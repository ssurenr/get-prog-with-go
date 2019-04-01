package main

import (
	"fmt"
	"math/rand"
)

const (
	width = 80
	height = 15
)

type Universe [][]bool

func main() {
	univ := NewUniverse()
	univ.Seed()
	univ.Show()
}

func NewUniverse() Universe {
	universe := make(Universe, height)

	for w := range universe {
		universe[w] = make([]bool, width)
	}

	return universe
}

func (u Universe) Show() {
	for h := range u {
		for w := range u[h] {
			if u[h][w] == true {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println("")
	}
}

func (u Universe) Seed() {
	for h := range u {
		for w:= range u[h] {
			if rand.Intn(3) == 0 {
				u[h][w] = true
			}
		}
	}
}