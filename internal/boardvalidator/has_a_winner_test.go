package boardvalidator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasWinner(t *testing.T) {
	testCases := []struct {
		name           string
		board          [3][3]string
		expectedOut    bool
		expectedWinner string
	}{
		{
			name: "empty board game ongoing",
			board: [3][3]string{
				{"", "", ""},
				{"", "", ""},
				{"", "", ""},
			},
			expectedOut:    false,
			expectedWinner: "",
		},
		{
			name: "board with turns realized game ongoing",
			board: [3][3]string{
				{"", "O", ""},
				{"", "X", ""},
				{"", "O", ""},
			},
			expectedOut:    false,
			expectedWinner: "",
		},
		{
			name: "board with player X won in first row",
			board: [3][3]string{
				{"X", "X", "X"},
				{"", "O", ""},
				{"", "O", ""},
			},
			expectedOut:    true,
			expectedWinner: "X",
		},
		{
			name: "board with player X won in second row",
			board: [3][3]string{
				{"", "O", ""},
				{"X", "X", "X"},
				{"", "O", ""},
			},
			expectedOut:    true,
			expectedWinner: "X",
		},
		{
			name: "board with player X won in third row",
			board: [3][3]string{
				{"", "O", ""},
				{"O", "", ""},
				{"X", "X", "X"},
			},
			expectedOut:    true,
			expectedWinner: "X",
		},
		{
			name: "board with player O won in first row",
			board: [3][3]string{
				{"O", "O", "O"},
				{"", "X", ""},
				{"", "X", ""},
			},
			expectedOut:    true,
			expectedWinner: "O",
		},
		{
			name: "board with player O won in second row",
			board: [3][3]string{
				{"", "O", ""},
				{"O", "O", "O"},
				{"", "X", ""},
			},
			expectedOut:    true,
			expectedWinner: "O",
		},
		{
			name: "board with player O won in third row",
			board: [3][3]string{
				{"", "X", ""},
				{"X", "", ""},
				{"O", "O", "O"},
			},
			expectedOut:    true,
			expectedWinner: "O",
		},
		{
			name: "board with player X won in first column",
			board: [3][3]string{
				{"X", "O", ""},
				{"X", "", ""},
				{"X", "O", "O"},
			},
			expectedOut:    true,
			expectedWinner: "X",
		},
		{
			name: "board with player X won in second column",
			board: [3][3]string{
				{"O", "X", ""},
				{"", "X", ""},
				{"O", "X", "O"},
			},
			expectedOut:    true,
			expectedWinner: "X",
		},
		{
			name: "board with player X won in third column",
			board: [3][3]string{
				{"O", "", "X"},
				{"", "", "X"},
				{"O", "O", "X"},
			},
			expectedOut:    true,
			expectedWinner: "X",
		},
		{
			name: "board with player O won in first column",
			board: [3][3]string{
				{"O", "X", ""},
				{"O", "", ""},
				{"O", "X", "X"},
			},
			expectedOut:    true,
			expectedWinner: "O",
		},
		{
			name: "board with player O won in second column",
			board: [3][3]string{
				{"X", "O", ""},
				{"", "O", ""},
				{"X", "O", "X"},
			},
			expectedOut:    true,
			expectedWinner: "O",
		},
		{
			name: "board with player O won in third column",
			board: [3][3]string{
				{"X", "", "O"},
				{"", "", "O"},
				{"X", "X", "O"},
			},
			expectedOut:    true,
			expectedWinner: "O",
		},
		{
			name: "board with player O won in diagonal",
			board: [3][3]string{
				{"O", "X", "X"},
				{"", "O", ""},
				{"X", "X", "O"},
			},
			expectedOut:    true,
			expectedWinner: "O",
		},
		{
			name: "board with player X won in diagonal",
			board: [3][3]string{
				{"X", "O", "O"},
				{"", "X", ""},
				{"O", "O", "X"},
			},
			expectedOut:    true,
			expectedWinner: "X",
		},
		{
			name: "board with player X won in secondary diagonal",
			board: [3][3]string{
				{"O", "O", "X"},
				{"", "X", ""},
				{"X", "O", "O"},
			},
			expectedOut:    true,
			expectedWinner: "X",
		},
		{
			name: "board with player O won in secondary diagonal",
			board: [3][3]string{
				{"X", "X", "O"},
				{"", "O", ""},
				{"O", "X", "X"},
			},
			expectedOut:    true,
			expectedWinner: "O",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			// Act
			out, outWinner := HasAWinner(&testCase.board)

			// Assert
			assert.Equal(t, testCase.expectedOut, out)
			assert.Equal(t, testCase.expectedWinner, outWinner)
		})
	}
}
