package playerturn

import "github.com/matheus-oliveira-andrade/tic-tac-toe/config/constants"

func GetPlayerTurn(playerTurn string) string {
	if playerTurn == "" {
		return constants.PlayerOne
	}

	if playerTurn == constants.PlayerOne {
		return constants.PlayerTwo
	}

	return constants.PlayerOne
}
