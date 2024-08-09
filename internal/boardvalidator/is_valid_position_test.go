package boardvalidator

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidPosition(t *testing.T) {
	testCases := []struct {
		name          string
		board         [3][3]string
		row           int
		column        int
		expectedOut   bool
		expectedError error
	}{
		{
			name: "row greater than board size",
			board: [3][3]string{
				{"X", "X", ""},
				{"X", "O", ""},
				{"O", "O", ""},
			},
			row:           10,
			column:        0,
			expectedOut:   false,
			expectedError: errors.New("invalid row"),
		},
		{
			name: "row less than board min size",
			board: [3][3]string{
				{"X", "X", ""},
				{"X", "O", ""},
				{"O", "O", ""},
			},
			row:           -1,
			column:        0,
			expectedOut:   false,
			expectedError: errors.New("invalid row"),
		},
		{
			name: "column greater than board min size",
			board: [3][3]string{
				{"X", "X", ""},
				{"X", "O", ""},
				{"O", "O", ""},
			},
			row:           0,
			column:        10,
			expectedOut:   false,
			expectedError: errors.New("invalid column"),
		},
		{
			name: "column less than board min size",
			board: [3][3]string{
				{"X", "X", ""},
				{"X", "O", ""},
				{"O", "O", ""},
			},
			row:           0,
			column:        -1,
			expectedOut:   false,
			expectedError: errors.New("invalid column"),
		},
		{
			name: "position in row and column already filled",
			board: [3][3]string{
				{"X", "X", ""},
				{"X", "O", ""},
				{"O", "O", ""},
			},
			row:           0,
			column:        0,
			expectedOut:   false,
			expectedError: errors.New("position already filled"),
		},
		{
			name: "valid position",
			board: [3][3]string{
				{"X", "X", ""},
				{"X", "O", ""},
				{"O", "O", ""},
			},
			row:           0,
			column:        2,
			expectedOut:   true,
			expectedError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			// Act
			out, err := IsValidPosition(&testCase.board, testCase.row, testCase.column)

			// Assert
			assert.Equal(t, testCase.expectedOut, out)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}
