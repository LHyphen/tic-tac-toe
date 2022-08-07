package main

import "fmt"

func main() {
	board := &board{
		square:    [][]Symbol{{Dot, Dot, Dot}, {Dot, Dot, Dot}, {Dot, Dot, Dot}},
		dimension: 3,
	}
	player1 := &humanPlayer{
		symbol: Cross,
		id:     1,
	}
	player2 := &humanPlayer{
		symbol: Circle,
		id:     2,
	}

	fmt.Println("\nThis is a Tic Tac Toe game")
	fmt.Println("The chessboard:")
	fmt.Printf("    y 0 1 2\n" +
		"  x +------\n" +
		"  0 | . . .\n" +
		"  1 | . . .\n" +
		"  2 | . . .\n\n")

	game := initGame(board, player1, player2)
	err := game.play()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	game.printResult()
}
