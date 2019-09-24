package main

import (
	"fmt"
	"math/rand"
)

const (
	width  = 80
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
		for w := range u[h] {
			if rand.Intn(3) == 0 {
				u[h][w] = true
			}
		}
	}
}

func (u Universe) Alive(x, y int) bool {
	x = (x + height) % height
	y = (y + width) % width
	return u[x][y]
}

func (u Universe) Neighbours(x, y int) int {
	neighbours := 0
	for h := x - 1; h <= x+1; h++ {
		for w := y - 1; w <= y+1; w++ {
			if h == x && w == y {
				continue
			} else {
				if u.Alive(h, w) {
					neighbours += 1
				}
			}
		}
	}
	fmt.Println(neighbours)
	return neighbours
}

func (u Universe) Next(x, y int) bool {
	neighbours := u.Neighbours(x, y)
	alive := false

	if u.Alive(x, y) {
		switch neighbours {
		case 0, 1:
			alive = false
		case 2, 3:
			alive = true
		default:
			alive = false
		}
	} else {
		if neighbours == 3 {
			alive = true
		}
	}

	return alive
}
