package boardvalidator

import "github.com/matheus-oliveira-andrade/tic-tac-toe/config/constants"

func HasAWinner(board *[3][3]string) (bool, string) {
	var success bool
	var player string

	if success, player = hasWinnerInRows(board); success {
		return true, player
	}

	if success, player = hasWinnerInColumns(board); success {
		return true, player
	}

	if success, player = hasInDiagonalsWinner(board); success {
		return true, player
	}

	return false, ""
}

func hasWinnerInRows(board *[3][3]string) (bool, string) {
	var lastViewedPlayer string
	var times int

	for r, a := range board {
		for c := range a {
			if lastViewedPlayer == "" {
				lastViewedPlayer = board[r][c]
				times = 1
				continue
			}

			if board[r][c] == lastViewedPlayer {
				times++
			} else {
				lastViewedPlayer = board[r][c]
				times = 1
			}

			if times == 3 {
				return true, lastViewedPlayer
			}
		}

		lastViewedPlayer = ""
		times = 0
	}

	return false, ""
}

func hasWinnerInColumns(board *[3][3]string) (bool, string) {
	for c := 0; c < 3; c++ {
		lastViewedPlayer := ""
		times := 0

		for r := 0; r < 3; r++ {
			if lastViewedPlayer == "" {
				lastViewedPlayer = board[r][c]
				times = 1
				continue
			}

			if board[r][c] == lastViewedPlayer {
				times++
			} else {
				lastViewedPlayer = board[r][c]
				times = 1
			}

			if times == 3 {
				return true, lastViewedPlayer
			}
		}
	}

	return false, ""
}

func hasInDiagonalsWinner(board *[3][3]string) (bool, string) {
	lastViewedPlayer := ""
	times := 0

	for rowIdx := range board {
		if board[rowIdx][rowIdx] == "" {
			continue
		}

		if lastViewedPlayer == "" {
			lastViewedPlayer = board[rowIdx][rowIdx]
			times++
			continue
		}

		if board[rowIdx][rowIdx] == lastViewedPlayer {
			times++
		} else {
			times = 0
			lastViewedPlayer = board[rowIdx][rowIdx]
		}

		if times == 3 {
			return true, lastViewedPlayer
		}

	}

	columnCounter := constants.BoardSize - 1
	lastViewedPlayer = ""
	times = 0

	for rowIdx := range board {
		if board[rowIdx][columnCounter] == "" {
			columnCounter--
			continue
		}

		if lastViewedPlayer == "" {
			lastViewedPlayer = board[rowIdx][columnCounter]
			times = 1
		} else if board[rowIdx][columnCounter] == lastViewedPlayer {
			times++
		} else {
			times = 1
			lastViewedPlayer = board[rowIdx][columnCounter]
		}

		if times == 3 {
			return true, lastViewedPlayer
		}

		columnCounter--
	}

	return false, ""
}
