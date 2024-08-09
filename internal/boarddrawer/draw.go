package boarddrawer

import "fmt"

func Draw(board *[3][3]string, playerTurn string, gameFinished bool) {
	fmt.Println("    0   1   2  ")
	fmt.Println("  +---+---+---+")

	for i, a := range board {
		fmt.Print(i, " ", "|")

		for j := range a {
			if board[i][j] == "" {
				fmt.Print(" ")
			}

			fmt.Print(" ", board[i][j], " |")
		}

		fmt.Print("\n")
		fmt.Println("  +---+---+---+")
	}

	if !gameFinished {
		fmt.Printf("Turn of player %v\n\n", playerTurn)
	}
}
