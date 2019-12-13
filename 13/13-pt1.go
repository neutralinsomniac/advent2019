package main

import (
	"fmt"
	"os"

	"github.com/neutralinsomniac/advent2019/intcode"
)

type TileID int

const (
	empty TileID = iota
	wall
	block
	paddle
	ball
)

type Vector struct {
	x, y int
}
type Tile struct {
	tileID  TileID
	heading Vector
}

type Coord struct {
	x, y int
}

type Board map[Coord]Tile

func main() {
	program := intcode.Program{}

	fmt.Println("*** PART 1 ***")
	program.InitStateFromFile(os.Args[1])

	halted := false

	board := make(Board)

	for !halted {
		var x, y int
		var tileID TileID
		var tmp int
		x, halted = program.RunUntilOutput()
		y, halted = program.RunUntilOutput()
		tmp, halted = program.RunUntilOutput()
		tileID = TileID(tmp)
		if !halted {
			tile := Tile{tileID: tileID}
			board[Coord{x, y}] = tile
		}
	}
	numBlockTiles := 0
	for _, t := range board {
		if t.tileID == block {
			numBlockTiles++
		}
	}
	fmt.Println(numBlockTiles)
}
