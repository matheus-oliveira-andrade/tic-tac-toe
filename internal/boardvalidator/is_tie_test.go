package boardvalidator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsTie(t *testing.T) {
	testCases := []struct {
		name        string
		board       [3][3]string
		expectedOut bool
	}{
		{
			name: "game ongoing",
			board: [3][3]string{
				{"X", "X", ""},
				{"X", "O", ""},
				{"O", "O", ""},
			},
			expectedOut: false,
		},
		{
			name: "tie game",
			board: [3][3]string{
				{"X", "X", "O"},
				{"O", "X", "X"},
				{"X", "O", "O"},
			},
			expectedOut: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			// Act
			out := IsTie(&testCase.board)

			// Assert
			assert.Equal(t, testCase.expectedOut, out)
		})
	}
}
