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

type Coord struct {
	x, y int
}

type Board map[Coord]TileID

type Game struct {
	board                  Board
	score                  int
	paddle                 Coord
	ball, prevball         Coord
	minX, minY, maxX, maxY int
}

func (g *Game) Draw() {
	fmt.Println("score:", g.score)
	for coord := range g.board {
		if coord.x > g.maxX {
			g.maxX = coord.x
		}
		if coord.x < g.minX {
			g.minX = coord.x
		}
		if coord.y > g.maxY {
			g.maxY = coord.y
		}
		if coord.y < g.minY {
			g.minY = coord.y
		}
	}
	for y := g.minY; y <= g.maxY; y++ {
		for x := g.minX; x <= g.maxX; x++ {
			for coord, tileID := range g.board {
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

	game := Game{}
	game.board = make(Board)
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
			paddlepos := game.paddle
			ballpos := game.ball
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
		case intcode.Output:
			switch output {
			case -1:
				// score
				_, halted = program.RunUntilOutput()
				game.score, halted = program.RunUntilOutput()
			default:
				// x, y, tile:
				x = output
				y, halted = program.RunUntilOutput()
				tmp, halted = program.RunUntilOutput()
				tileID = TileID(tmp)
				game.board[Coord{x, y}] = tileID
				if tileID == ball {
					game.ball = Coord{x, y}
					if draw && game.ball != game.prevball {
						game.Draw()
						game.prevball = game.ball
					}
				} else if tileID == paddle {
					game.paddle = Coord{x, y}
				}
			}
		}
	}
	fmt.Println("score:", game.score)
}
