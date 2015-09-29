package main

import (
	"wumpus"
)

type Map struct {
	Name          string
	Width, Height uint8
	Arr           [][]rune
}

func main() {
	gb := wumpus.NewGameBot()
	gb.Run()
}
