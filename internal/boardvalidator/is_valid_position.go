package boardvalidator

import (
	"errors"
	"github.com/matheus-oliveira-andrade/tic-tac-toe/config/constants"
)

func IsValidPosition(board *[3][3]string, row int, column int) (bool, error) {
	if row > constants.BoardSize-1 || row < 0 {
		return false, errors.New("invalid row")
	}

	if column > constants.BoardSize-1 || column < 0 {
		return false, errors.New("invalid column")
	}

	if board[row][column] != "" {
		return false, errors.New("position already filled")
	}

	return true, nil
}
