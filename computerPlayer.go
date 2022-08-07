package main

import "fmt"


type computerPlayer struct {
	symbol Symbol
	id     int
}

func (c *computerPlayer) getSymbol() Symbol {
	return c.symbol
}

func (c *computerPlayer) getNextMove() (int, int, error) {
	// to be implemented
	return 0, 0, fmt.Errorf("Invalid Symbol")
}

func (c *computerPlayer) getId() int {
	return c.id
}
