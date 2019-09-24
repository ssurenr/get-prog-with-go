package main

import (
	"fmt"
)

func main() {
	var board [8][8]rune

	board[0] = [8]rune{'n', 'r', 'b', 'k', 'q', 'b', 'r', 'n'}
	for column := range board[1] {
		board[1][column] = 'p'
		board[6][column] = 'P'
	}
	board[7] = [8]rune{'N', 'R', 'B', 'K', 'Q', 'B', 'R', 'N'}
	printBoard(board)
}

func printBoard(board [8][8]rune) {
	for _, row := range board {
		for _, item := range row {
			fmt.Printf("%c", item)
		}
		fmt.Println()
	}
}
