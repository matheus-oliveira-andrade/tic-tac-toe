package winnerdrawer

import "fmt"

func Draw(playerTurn string) {
	fmt.Println("+----------------------+")
	fmt.Println("|                      |")
	fmt.Printf("|   ** WINNER: %v **    |\n", playerTurn)
	fmt.Println("|                      |")
	fmt.Println("+----------------------+")
}
