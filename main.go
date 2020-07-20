package main

import (
	"strconv"
	"fmt"
	"strings"
	"math"
)

func main() {
	// Creating new game
	g := newGame()
	setInitialGameData(g)
	play(g)
	checkFinalState(g)

	fmt.Println("Thank you for playing!")
}

type game struct {
	board [9]string
	userTurn bool
	userSymbol string
	computerSymbol string
}

func newGame() *game {
	g := &game{}
	for i := 0; i < 9; i++ {
		g.board[i] = strconv.Itoa(i+1)
	}
	return g
}

func setInitialGameData(g *game) {
	// Determine user's symbol
	fmt.Println("X or O?")
	var chosenSymbol string
	fmt.Scanf("%s\n", &chosenSymbol)
	g.userSymbol = strings.ToUpper(chosenSymbol)

	// Set symbols
	if g.userSymbol == "X" {
		g.computerSymbol = "O"
	} else {
		g.computerSymbol = "X"
	}

	// Set first turn
	g.userTurn = "X" == g.userSymbol
}

func play(g *game) {
	fmt.Println("\nUse the positions below to play:")
	outputBoard(g.board)

	// Loop until win or tie is detected
	for !hasWinner(g.board) && !hasTie(g.board) {
		if g.userTurn {
			// Ask user for position
			var choice int
			fmt.Println("Choose an open position:")
			fmt.Scanf("%d\n", &choice)

			// If the position isn't valid, try again
			if !checkValidMove(choice, g.board) {
				fmt.Println("Invalid position, try again.")
				continue;
			}

			// Add position to board
			g.board[choice-1] = g.userSymbol
		} else {
			// Ask bot to make selection
			makeBotSelection(g)
		}

		g.userTurn = !g.userTurn
		outputBoard(g.board)
	}
}

func outputBoard(board [9]string) {
	fmt.Println("---------")

	offset := 0
	for i := 0; i < 3; i++ {
		fmt.Print("- ")
		for j := 0; j < 3; j++ {
			fmt.Print(board[j + offset] + " ")
		}
		fmt.Println("-")
		offset += 3
	}

	fmt.Println("---------")
}

func makeBotSelection(g *game) {
	bestScore := -math.MaxUint32
	var score int
	var move int

	for i := 0; i < 9; i++ {
		if g.board[i] != "X" && g.board[i] != "O" {
			game := g
			game.board[i] = g.computerSymbol
			game.userTurn = !game.userTurn
			score = minimax(game, 0, false)
			game.board[i] = strconv.Itoa(i+1)
			game.userTurn = !game.userTurn

			if score > bestScore {
				bestScore = score
				move = i
			}
		}
	}

	g.board[move] = g.computerSymbol
}

func minimax(game *game, depth int, userTurn bool) int {
	if hasWinner(game.board) {
		if game.userTurn {
			return 1
		} else {
			return -1
		}
	}

	if hasTie(game.board) {
		return 0
	}

	if userTurn {
		bestScore := -math.MaxUint32
		for i := 0; i < 9; i++ {
			if game.board[i] != "X" && game.board[i] != "O" {
				game.board[i] = game.computerSymbol
				game.userTurn = !game.userTurn
				score := minimax(game, depth+1, false)
				game.board[i] = strconv.Itoa(i+1)
				game.userTurn = !game.userTurn

				if score > bestScore {
					bestScore = score
				}
			}
		}
		return bestScore
	} else {
		bestScore := math.MaxUint32
		for i := 0; i < 9; i++ {
			if game.board[i] != "X" && game.board[i] != "O" {
				game.board[i] = game.userSymbol
				game.userTurn = !game.userTurn
				score := minimax(game, depth+1, true)
				game.board[i] = strconv.Itoa(i+1)
				game.userTurn = !game.userTurn

				if score < bestScore {
					bestScore = score
				}
			}
		}
		return bestScore
	}
}

func checkValidMove(choice int, board [9]string) bool {
	if choice < 1 || choice > 9 {
		return false
	} else if _, err := strconv.Atoi(board[choice-1]); err != nil {
		return false
	} else {
		return true
	}
}

func checkFinalState(g *game) {
	if hasWinner(g.board) {
		var winner string
		if g.userTurn {
			winner = g.computerSymbol
		} else {
			winner = g.userSymbol
		}

		fmt.Println(fmt.Sprintf("Winner: %s", winner))
	} else {
		fmt.Println("It's a tie!")
	}
}

func hasWinner(b [9]string) bool {
	winningCombinations := [8][3]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}

	for _, combo := range winningCombinations {
		if b[combo[0]] == b[combo[1]] && b[combo[1]] == b[combo[2]] {
			return true
		}
	}

	return false
}

func hasTie(b [9]string) bool {
	for _, position := range b {
		if position != "X" && position != "O" {
			return false
		}
	}

	return true
}