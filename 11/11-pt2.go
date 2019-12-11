package main

import (
	"fmt"
	"math"
	"os"

	"github.com/neutralinsomniac/advent2019/intcode"
)

type Coord struct {
	x, y int
}

type Tile struct {
	color   int
	painted bool
}

type Hull map[Coord]Tile

func (h *Hull) PaintTile(coord Coord, color int) {
	newTile := Tile{painted: true, color: color}
	(*h)[coord] = newTile
}

type Heading int

const (
	up Heading = iota
	right
	down
	left
)

type Direction int

const (
	turnLeft  Direction = 0
	turnRight           = 1
)

const (
	black = 0
	white = 1
)

type Robot struct {
	heading Heading
	pos     Coord
}

func (r *Robot) Turn(d Direction) {
	switch d {
	case turnLeft:
		if r.heading == up {
			r.heading = left
		} else {
			r.heading -= 1
		}
	case turnRight:
		r.heading = (r.heading + 1) % 4
	default:
		panic(fmt.Sprintf("wtf kind of direction is this: %d\n", d))
	}
}

func (r *Robot) MoveForward() {
	switch r.heading {
	case up:
		r.pos.y--
	case down:
		r.pos.y++
	case left:
		r.pos.x--
	case right:
		r.pos.x++
	}
}

func (r *Robot) GetPos() Coord {
	return r.pos
}

func main() {
	program := intcode.Program{}

	fmt.Println("*** PART 2 ***")
	program.InitStateFromFile(os.Args[1])

	halted := false

	hull := make(Hull)
	robot := Robot{heading: up}
	// direct reference since the initial tile technically isn't painted; it's just white already
	hull[robot.pos] = Tile{color: white, painted: false}

	for !halted {
		colorUnderRobot := hull[robot.GetPos()].color
		program.SetReaderFromInts(colorUnderRobot)
		var colorToPaint, directionToTurn int
		colorToPaint, halted = program.RunUntilOutput()
		program.SetReader(nil)
		directionToTurn, halted = program.RunUntilOutput()
		if !halted {
			hull.PaintTile(robot.pos, colorToPaint)
			robot.Turn(Direction(directionToTurn))
			robot.MoveForward()
		}
	}

	// first figure out our world boundaries
	minX := math.MaxInt32
	minY := math.MaxInt32
	maxX := -1 * math.MaxInt32
	maxY := -1 * math.MaxInt32
	for coord, _ := range hull {
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
			if hull[Coord{x, y}].color == white && hull[Coord{x, y}].painted {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}
