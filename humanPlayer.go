package main

import "fmt"

var (
// MovePlayer1 = [4][2]int{{1, 1}, {2, 0}, {2, 2}, {2, 1}}
// MovePlayer2 = [4][2]int{{1, 2}, {0, 2}, {0, 0}, {0, 0}}
)

type humanPlayer struct {
	symbol Symbol
	index  int
	id     int
}

func (h *humanPlayer) getSymbol() Symbol {
	return h.symbol
}

// func (h *humanPlayer) getNextMove() (int, int, error) {
// 	if h.symbol == Cross {
// 		h.index = h.index + 1
// 		return MovePlayer1[h.index-1][0], MovePlayer1[h.index-1][1], nil
// 	} else if h.symbol == Circle {
// 		h.index = h.index + 1
// 		return MovePlayer2[h.index-1][0], MovePlayer2[h.index-1][1], nil
// 	}
// 	return 0, 0, fmt.Errorf("Invalid Symbol")
// }

func (h *humanPlayer) getNextMove() (int, int, error) {
	symbolString := ""
	if h.symbol == Cross {
		symbolString = "x"
	} else if h.symbol == Circle {
		symbolString = "o"
	} else {
		return 0, 0, fmt.Errorf("Invalid Symbol")
	}

	x, y := -1, -1
	fmt.Printf("Player %d Please enter the position of the symbol %s\n", h.id, symbolString)
	fmt.Printf("X: ")
	fmt.Scanln(&x)
	fmt.Printf("Y: ")
	fmt.Scanln(&y)

	h.index = h.index + 1
	return x, y, nil
}

func (h *humanPlayer) getId() int {
	return h.id
}
