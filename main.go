package main

import (
	"fmt"
	"math/rand"
	"time"
)

var l = fmt.Println

func main() {
	gridDefinitions := initialize()
	FirstBoard, NextBoard := startBoards(gridDefinitions)
	FirstBoard.Seed()
	for {
		FirstBoard.Display()
		Step(*FirstBoard, *NextBoard)
		FirstBoard, NextBoard = NextBoard, FirstBoard
		time.Sleep(time.Duration(gridDefinitions.SleepIteration) * time.Millisecond)
		l(gridDefinitions.AnsiEscapeSeq)
	}
}

func startBoards(gridDefinitions GridDefinitions) (*Board, *Board) {
	return MakeWBoard(gridDefinitions), MakeWBoard(gridDefinitions)
}

func initialize() GridDefinitions {
	g := Setup()
	l(g.AnsiEscapeSeq)
	rand.Seed(time.Now().UTC().UnixNano())
	return g
}
