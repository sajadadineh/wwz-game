package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

type Location struct {
	x int
	y int
}

type Bullet struct {
	location Location
}

type Player struct {
	symbol   rune
	location Location
	died     bool
}

type Zambi struct {
	symbol   rune
	location Location
	died     bool
}

type World struct {
	player  Player
	height  int
	width   int
	bullets []Bullet
	zambi   []Zambi
}

func main() {
	maxX, maxY := termbox.Size()
	fmt.Println(maxX, maxY)
}
