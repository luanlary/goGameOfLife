package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	gridDefinitions := Setup()
	fmt.Println(gridDefinitions.AnsiEscapeSeq)
	rand.Seed(time.Now().UTC().UnixNano())
	newWorld := MakeWorld(gridDefinitions)
	nextWorld := MakeWorld(gridDefinitions)
	newWorld.Seed()
	for {
		newWorld.Display()
		Step(*newWorld, *nextWorld)
		newWorld, nextWorld = nextWorld, newWorld
		time.Sleep(time.Duration(gridDefinitions.SleepIteration) * time.Millisecond)
		fmt.Println(gridDefinitions.AnsiEscapeSeq)
	}
}
