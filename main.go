package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed once

	var matrix [3][3]string
	// Initialize the board with spaces
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = " "
		}
	}

	var row, col int
	var val string

	for {
		// Player move
		fmt.Print("Enter row (0-2), column (0-2) and value as X: ")
		fmt.Scan(&row, &col, &val)

		if isValidInput(row, col, val, matrix) {
			matrix[row][col] = "X"
			printBoard(matrix)
		} else {
			fmt.Println("Invalid input or cell already taken. Try again.")
			continue
		}

		if winner, line := checkWin(matrix); winner != "" {
			fmt.Printf("Player %s wins on %s!\n", winner, line)
			break
		}

		// AI Move
		row, col = bestMove(matrix, "O", "X")
		if row != -1 {
			matrix[row][col] = "O"
			fmt.Println("Computer played:")
			printBoard(matrix)
		}

		if winner, line := checkWin(matrix); winner != "" {
			fmt.Printf("Player %s wins on %s!\n", winner, line)
			break
		}

		// Check for draw
		if isBoardFull(matrix) {
			fmt.Println("It's a draw! Board is full.")
			break
		}
	}
}

func isValidInput(row, col int, val string, board [3][3]string) bool {
	return row >= 0 && row < 3 &&
		col >= 0 && col < 3 &&
		strings.ToUpper(val) == "X" &&
		board[row][col] == " "
}

func checkWin(board [3][3]string) (string, string) {
	winLines := []struct {
		coords [3][2]int
		name   string
	}{
		{coords: [3][2]int{{0, 0}, {0, 1}, {0, 2}}, name: "Top Row"},
		{coords: [3][2]int{{1, 0}, {1, 1}, {1, 2}}, name: "Middle Row"},
		{coords: [3][2]int{{2, 0}, {2, 1}, {2, 2}}, name: "Bottom Row"},
		{coords: [3][2]int{{0, 0}, {1, 0}, {2, 0}}, name: "Left Column"},
		{coords: [3][2]int{{0, 1}, {1, 1}, {2, 1}}, name: "Middle Column"},
		{coords: [3][2]int{{0, 2}, {1, 2}, {2, 2}}, name: "Right Column"},
		{coords: [3][2]int{{0, 0}, {1, 1}, {2, 2}}, name: "Main Diagonal"},
		{coords: [3][2]int{{0, 2}, {1, 1}, {2, 0}}, name: "Anti-Diagonal"},
	}

	for _, line := range winLines {
		r1, c1 := line.coords[0][0], line.coords[0][1]
		r2, c2 := line.coords[1][0], line.coords[1][1]
		r3, c3 := line.coords[2][0], line.coords[2][1]

		if board[r1][c1] != " " &&
			board[r1][c1] == board[r2][c2] &&
			board[r2][c2] == board[r3][c3] {
			return board[r1][c1], line.name
		}
	}
	return "", ""
}

func printBoard(board [3][3]string) {
	fmt.Println("Tic-Tac-Toe Board:")
	for i := 0; i < 3; i++ {
		fmt.Printf(" %s | %s | %s \n", board[i][0], board[i][1], board[i][2])
		if i < 2 {
			fmt.Println("---+---+---")
		}
	}
}

func randomCoordinate() (int, int) {
	return rand.Intn(3), rand.Intn(3)
}

func isBoardFull(board [3][3]string) bool {
	for _, row := range board {
		for _, val := range row {
			if val == " " {
				return false
			}
		}
	}
	return true
}

func bestMove(board [3][3]string, aiSymbol, playerSymbol string) (int, int) {
	// 1. Check for winning move
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == " " {
				board[i][j] = aiSymbol
				if winner, _ := checkWin(board); winner == aiSymbol {
					board[i][j] = " " // reset
					return i, j
				}
				board[i][j] = " "
			}
		}
	}

	// 2. Block player's winning move
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == " " {
				board[i][j] = playerSymbol
				if winner, _ := checkWin(board); winner == playerSymbol {
					board[i][j] = " "
					return i, j
				}
				board[i][j] = " "
			}
		}
	}

	// 3. Take center
	if board[1][1] == " " {
		return 1, 1
	}

	// 4. Take corners
	corners := [][2]int{{0, 0}, {0, 2}, {2, 0}, {2, 2}}
	for _, c := range corners {
		if board[c[0]][c[1]] == " " {
			return c[0], c[1]
		}
	}

	// 5. Take sides
	sides := [][2]int{{0, 1}, {1, 0}, {1, 2}, {2, 1}}
	for _, s := range sides {
		if board[s[0]][s[1]] == " " {
			return s[0], s[1]
		}
	}

	// Fallback: no move found
	return -1, -1
}
