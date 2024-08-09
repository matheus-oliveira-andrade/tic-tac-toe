package boardvalidator

func IsTie(board *[3][3]string) bool {
	for i, a := range board {
		for j := range a {
			if board[i][j] == "" {
				return false
			}
		}
	}

	return true
}
