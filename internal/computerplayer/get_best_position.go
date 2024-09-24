package computerplayer

import (
	"math"

	"github.com/matheus-oliveira-andrade/tic-tac-toe/config/constants"
	"github.com/matheus-oliveira-andrade/tic-tac-toe/internal/boardvalidator"
)

func GetBestPosition(board *[3][3]string) (int, int) {
	bestScore := math.MinInt
	var bestRow, bestColumn int

	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if board[r][c] != "" {
				continue
			}

			board[r][c] = constants.Computer
			score := miniMax(board, 0, false)
			board[r][c] = ""

			if score > bestScore {
				bestScore = score
				bestRow = r
				bestColumn = c
			}
		}
	}

	return bestRow, bestColumn
}

var scores = map[string]int{
	constants.PlayerOne: -10,
	constants.Computer:  10,
	"tie":               0,
}

func miniMax(board *[3][3]string, depth int, isMaximizing bool) int {
	if hasWinner, player := boardvalidator.HasAWinner(board); hasWinner {
		return scores[player]
	}

	if boardvalidator.IsTie(board) {
		return scores["tie"]
	}

	if isMaximizing {
		bestScore := math.MinInt

		for r := 0; r < 3; r++ {
			for c := 0; c < 3; c++ {
				if board[r][c] != "" {
					continue
				}

				board[r][c] = constants.Computer
				score := miniMax(board, depth+1, false)
				board[r][c] = ""

				if score > bestScore {
					bestScore = score
				}
			}
		}

		return bestScore
	} else {
		bestScore := math.MaxInt

		for r := 0; r < 3; r++ {
			for c := 0; c < 3; c++ {
				if board[r][c] != "" {
					continue
				}

				board[r][c] = constants.PlayerOne
				score := miniMax(board, depth+1, true)
				board[r][c] = ""

				if score < bestScore {
					bestScore = score
				}
			}
		}

		return bestScore
	}
}
