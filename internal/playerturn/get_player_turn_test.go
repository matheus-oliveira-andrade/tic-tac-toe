package playerturn

import (
	"github.com/matheus-oliveira-andrade/tic-tac-toe/config/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPlayerTurn(t *testing.T) {
	testCases := []struct {
		name               string
		playerTurn         string
		expectedPlayerTurn string
	}{
		{
			name:               "given empty last turn should return turn X",
			playerTurn:         "",
			expectedPlayerTurn: constants.PlayerOne,
		},
		{
			name:               "given X as last turn should return turn of O",
			playerTurn:         constants.PlayerOne,
			expectedPlayerTurn: constants.PlayerTwo,
		},
		{
			name:               "given O as last turn should return turn of X",
			playerTurn:         constants.PlayerTwo,
			expectedPlayerTurn: constants.PlayerOne,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			// Act
			playerTurn := GetPlayerTurn(testCase.playerTurn)

			// Assert
			assert.Equal(t, testCase.expectedPlayerTurn, playerTurn)
		})
	}
}
