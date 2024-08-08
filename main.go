package main

import (
	"fmt"
	"github.com/matheus-oliveira-andrade/tic-tac-toe/config/constants"
	"github.com/matheus-oliveira-andrade/tic-tac-toe/internal/boarddrawer"
	"github.com/matheus-oliveira-andrade/tic-tac-toe/internal/boardvalidator"
	"github.com/matheus-oliveira-andrade/tic-tac-toe/internal/inputreader"
	"github.com/matheus-oliveira-andrade/tic-tac-toe/internal/playerturn"
	"github.com/matheus-oliveira-andrade/tic-tac-toe/internal/winnerdrawer"
	"github.com/matheus-oliveira-andrade/tic-tac-toe/utils"
	"time"
)

func main() {
	board := [constants.BoardSize][constants.BoardSize]string{}
	var playerTurn string

	for {
		if boardvalidator.HasAWinner(&board) {
			boarddrawer.Draw(&board, playerTurn, true)
			winnerdrawer.Draw(playerTurn)
			break
		}

		if boardvalidator.IsTie(&board) {
			fmt.Printf("Tie!!")
			break
		}

		playerTurn = playerturn.GetPlayerTurn(playerTurn)

		boarddrawer.Draw(&board, playerTurn, false)

		row, column := RequestPosition(&board)

		board[row][column] = playerTurn

		fmt.Printf("\nplayer %v assigned in position (%v)\n", playerTurn, fmt.Sprintf("%v, %v", row, column))

		time.Sleep(1 * time.Second)
		utils.ClearTerminal()
	}
}

func RequestPosition(board *[3][3]string) (int, int) {
	var row int
	var column int
	var err error

	for {
		row, column, err = inputreader.ReadPosition()
		if err != nil {
			fmt.Println(err)
			continue
		}

		_, err = boardvalidator.IsValidPosition(board, row, column)
		if err != nil {
			fmt.Println(err)
			continue
		}

		break
	}

	return row, column
}
