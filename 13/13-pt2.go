package main

import (
	"fmt"
	"os"

	"github.com/neutralinsomniac/advent2019/intcode"
)

// fuck you @jrick :)
var score int
var minX, minY, maxX, maxY int

type TileID int

const (
	empty TileID = iota
	wall
	block
	paddle
	ball
)

type Coord struct {
	x, y int
}

type Board map[Coord]TileID

func (b Board) GetPaddlePos() Coord {
	for coord, tileID := range b {
		if tileID == paddle {
			return coord
		}
	}
	return Coord{}
}

func (b Board) GetBallPos() Coord {
	for coord, tileID := range b {
		if tileID == ball {
			return coord
		}
	}
	return Coord{}
}

func (b Board) Draw() {
	for coord := range b {
		if coord.x > maxX {
			maxX = coord.x
		}
		if coord.x < minX {
			minX = coord.x
		}
		if coord.y > maxY {
			maxY = coord.y
		}
		if coord.y < minY {
			minY = coord.y
		}
	}
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			for coord, tileID := range b {
				if coord.x == x && coord.y == y {
					switch tileID {
					case wall:
						fmt.Printf("#")
					case block:
						fmt.Printf("X")
					case paddle:
						fmt.Printf("=")
					case ball:
						fmt.Printf("O")
					case empty:
						fmt.Printf(" ")
					}
				}
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	draw := false
	program := intcode.Program{}

	fmt.Println("*** PART 2 ***")
	program.InitStateFromFile(os.Args[1])

	// initiate free play
	program.SetMemory(0, 2)
	halted := false

	board := make(Board)
	if draw {
		fmt.Printf("\033[2J;\033[H")
	}

	for !halted {
		var x, y int
		var tileID TileID
		var tmp int
		var opcode intcode.Opcode
		var output int
		if draw {
			fmt.Printf("\033[H")
		}

		opcode, output, halted = program.RunUntilInputOrOutput()
		switch opcode {
		case intcode.Input:
			// find the paddle
			paddlepos := board.GetPaddlePos()
			// find the ball
			ballpos := board.GetBallPos()
			// move the joystick
			if paddlepos.x > ballpos.x {
				program.SetReaderFromInts(-1)
			} else if paddlepos.x < ballpos.x {
				program.SetReaderFromInts(1)
			} else {
				program.SetReaderFromInts(0)
			}
			// and run the input instruction
			program.Step()
			if draw {
				board.Draw()
				fmt.Println("score:", score)
			}
		case intcode.Output:
			switch output {
			case -1:
				// score
				_, halted = program.RunUntilOutput()
				score, halted = program.RunUntilOutput()
			default:
				// x, y, tile:
				x = output
				y, halted = program.RunUntilOutput()
				tmp, halted = program.RunUntilOutput()
				tileID = TileID(tmp)
				board[Coord{x, y}] = tileID
			}
		}
	}
	fmt.Println("score:", score)
}
