package main

import (
	"fmt"
	"github.com/matheus-oliveira-andrade/tic-tac-toe/config/constants"
	"github.com/matheus-oliveira-andrade/tic-tac-toe/internal/boarddrawer"
	"github.com/matheus-oliveira-andrade/tic-tac-toe/internal/boardvalidator"
	"github.com/matheus-oliveira-andrade/tic-tac-toe/internal/computerplayer"
	"github.com/matheus-oliveira-andrade/tic-tac-toe/internal/inputreader"
	"github.com/matheus-oliveira-andrade/tic-tac-toe/internal/playerturn"
	"github.com/matheus-oliveira-andrade/tic-tac-toe/internal/winnerdrawer"
	"github.com/matheus-oliveira-andrade/tic-tac-toe/utils"
	"time"
)

func CreateBoard() [constants.BoardSize][constants.BoardSize]string {
	return [constants.BoardSize][constants.BoardSize]string{}
}

func main() {
	withComputerPlayer := ReadPlayWithComputerPlayer()

	utils.ClearTerminal()

	ticTacToe(CreateBoard(), withComputerPlayer)
}

func ticTacToe(board [3][3]string, playWithComputer bool) {
	var playerTurn string

	for {
		if hasWinner, _ := boardvalidator.HasAWinner(&board); hasWinner {
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

		var row, column int

		if playWithComputer && playerTurn == constants.PlayerTwo {
			fmt.Printf("\n%v playing....\n", playerTurn)
			time.Sleep(1 * time.Second)

			row, column = computerplayer.GetBestPosition(&board)
		} else {
			row, column = RequestPosition(&board)
		}

		board[row][column] = playerTurn

		fmt.Printf("\nplayer %v assigned in position (%v)\n", playerTurn, fmt.Sprintf("%v, %v", row, column))

		time.Sleep(1 * time.Second)
		utils.ClearTerminal()
	}
}

func RequestPosition(board *[3][3]string) (int, int) {
	var row, column int
	var err error

	fmt.Printf("Type a position (e.a: 1,1): ")

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

func ReadPlayWithComputerPlayer() bool {
	fmt.Printf("Would you like to play against the computer? (y/n): ")

	for {
		answer, err := inputreader.ReadYesOrNot()
		if err != nil {
			fmt.Println(err)
			continue
		}

		return answer
	}
}
