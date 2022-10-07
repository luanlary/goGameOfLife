package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width          = 40
	height         = 20
	sleepIteration = 200
	ansiEscapeSeq  = "\033c\x0c"
	brownSquare    = "\xF0\x9F\x9F\xAB"
	greenSquare    = "\xF0\x9F\x9F\xA9"
)

type World [][]bool

func MakeWorld() World {
	w := make(World, height)
	for i := range w {
		w[i] = make([]bool, width)
	}
	return w
}

func (w World) Seed() {
	for _, row := range w {
		for i := range row {
			if rand.Intn(4) == 1 {
				row[i] = true
			}
		}
	}
}

func (w World) Display() {
	for _, row := range w {
		for _, cell := range row {
			switch {
			case cell:
				fmt.Printf(greenSquare)
			default:
				fmt.Printf(brownSquare)
			}
		}
		fmt.Printf("\n")
	}
}

func (w World) Alive(x, y int) bool {
	y = (height + y) % height
	x = (width + x) % width
	return w[y][x]
}

func (w World) Neighbors(x, y int) int {
	var neighbors int
	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if i == y && j == x {
				continue
			}
			if w.Alive(j, i) {
				neighbors++
			}
		}
	}
	return neighbors
}

func (w World) Next(x, y int) bool {
	n := w.Neighbors(x, y)
	alive := w.Alive(x, y)
	if n < 4 && n > 1 && alive {
		return true
	} else if n == 3 && !alive {
		return true
	} else {
		return false
	}
}

func Step(a, b World) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			b[i][j] = a.Next(j, i)
		}
	}
}

func main() {
	fmt.Println(ansiEscapeSeq)
	rand.Seed(time.Now().UTC().UnixNano())
	newWorld := MakeWorld()
	nextWorld := MakeWorld()
	newWorld.Seed()
	for {
		newWorld.Display()
		Step(newWorld, nextWorld)
		newWorld, nextWorld = nextWorld, newWorld
		time.Sleep(sleepIteration * time.Millisecond)
		fmt.Println(ansiEscapeSeq)
	}
}
