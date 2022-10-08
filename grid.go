package main

import (
	"fmt"
	"math/rand"
)

type Board struct {
	Grid            Grid
	GridDefinitions GridDefinitions
}

type Grid [][]bool

var p = fmt.Printf

func MakeWBoard(d GridDefinitions) *Board {
	board := new(Board)
	w := make(Grid, d.Height)
	for i := range w {
		w[i] = make([]bool, d.Width)
	}
	board.GridDefinitions = d
	board.Grid = w
	return board
}

func (b Board) Seed() {
	for _, row := range b.Grid {
		for i := range row {
			if rand.Intn(4) == 1 {
				row[i] = true
			}
		}
	}
}

func (b Board) Display() {
	for _, row := range b.Grid {
		for _, cell := range row {
			switch {
			case cell:
				p(b.GridDefinitions.BlueBlocks)
			default:
				p(b.GridDefinitions.RedBlocks)
			}
		}
		p("\n")
	}
}

func (b Board) Alive(x, y int) bool {
	y = (b.GridDefinitions.Height + y) % b.GridDefinitions.Height
	x = (b.GridDefinitions.Width + x) % b.GridDefinitions.Width
	return b.Grid[y][x]
}

func (b Board) Neighbors(x, y int) int {
	var neighbors int
	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if i == y && j == x {
				continue
			}
			if b.Alive(j, i) {
				neighbors++
			}
		}
	}
	return neighbors
}

func (b Board) Next(x, y int) bool {
	n := b.Neighbors(x, y)
	alive := b.Alive(x, y)
	if n < 4 && n > 1 && alive {
		return true
	} else if n == 3 && !alive {
		return true
	} else {
		return false
	}
}

func Step(a, b Board) {
	for i := 0; i < b.GridDefinitions.Height; i++ {
		for j := 0; j < b.GridDefinitions.Width; j++ {
			b.Grid[i][j] = a.Next(j, i)
		}
	}
}
