package main

import (
	"fmt"
	"math/rand/v2"
)

const header string = "+----------------+\n| Go Tic-Tac-Toe |\n+----------------+\n"

var grid = [][]string{
	{"_", "_", "_"},
	{"_", "_", "_"},
	{"_", "_", "_"},
}

var player1 bool = true
var max_turns int8 = 9

func turn() {
	var player rune
	srow := rand.IntN(3)
	scol := rand.IntN(3)

	for i, row := range grid {
		for j, col := range row {
			if i == srow && j == scol && col == "_" {
				if player1 {
					player = 'X'
					player1 = false
				} else {
					player = 'O'
					player1 = true
				}
				grid[i][j] = string(player)
				fmt.Printf("(%v, %s), %v\n", j+1, num_to_let(i), string(player))
				max_turns -= 1
			}
		}
	}
}

func won() (bool, string) {
	pos1 := grid[0][0]
	pos2 := grid[0][1]
	pos3 := grid[0][2]
	pos4 := grid[1][0]
	pos5 := grid[1][1]
	pos6 := grid[1][2]
	pos7 := grid[2][0]
	pos8 := grid[2][1]
	pos9 := grid[2][2]
	switch {
	case pos1 != "_" && pos1 == pos2 && pos2 == pos3:
		return true, string(pos1)
	case pos4 != "_" && pos4 == pos5 && pos5 == pos6:
		return true, string(pos4)
	case pos7 != "_" && pos7 == pos8 && pos8 == pos9:
		return true, string(pos7)
	case pos1 != "_" && pos1 == pos4 && pos4 == pos7:
		return true, string(pos1)
	case pos2 != "_" && pos2 == pos5 && pos5 == pos8:
		return true, string(pos2)
	case pos3 != "_" && pos3 == pos6 && pos6 == pos9:
		return true, string(pos3)
	case pos1 != "_" && pos1 == pos5 && pos5 == pos9:
		return true, string(pos1)
	case pos3 != "_" && pos3 == pos5 && pos5 == pos7:
		return true, string(pos3)
	case max_turns == 0:
		return true, "Tie"
	default:
		return false, ""
	}
}

func num_to_let(num int) string {
	switch num {
	case 0:
		return "A"
	case 1:
		return "B"
	case 2:
		return "C"
	default:
		return ""
	}
}

func print_board() {
	fmt.Printf("\n   1 2 3\n")
	for i := range len(grid) {
		fmt.Printf("%s %s\n", num_to_let(i), grid[i])
	}
}

func main() {
	fmt.Printf(header)
	fmt.Printf("\nGame Log:\n")

	is_over, winner := won()

	for !is_over {
		switch is_over {
		case false:
			turn()
			is_over, winner = won()
		}
	}

	switch winner {
	case "Tie":
		fmt.Printf("It's a %s\n", winner)
	case "X":
		fmt.Printf("Player %s has won!\n", winner)
	case "O":
		fmt.Printf("Player %s has won!\n", winner)
	}

	print_board()
}
