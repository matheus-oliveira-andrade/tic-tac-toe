package inputreader

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func ReadPosition() (int, int, error) {
	fmt.Printf("Type a position (e.a: 1,1): ")
	position := ""
	_, err := fmt.Scanln(&position)
	if err != nil {
		return 0, 0, errors.New("empty position is not allowed")
	}

	parts := strings.Split(position, ",")

	if len(parts) != 2 {
		return 0, 0, errors.New("invalid format of position")
	}

	row, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, errors.New("error parsing row at given position")
	}

	column, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, errors.New("error parsing column at given position")
	}

	return row, column, nil
}

func ReadYesOrNot() (bool, error) {
	answer := ""
	_, err := fmt.Scanln(&answer)
	if err != nil {
		return false, errors.New("invalid answer")
	}

	if answer == "n" || answer == "N" {
		return false, nil
	}

	if answer == "y" || answer == "Y" {
		return true, nil
	}

	return false, errors.New("invalid answer")
}
