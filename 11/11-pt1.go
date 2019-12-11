package main

import (
	"fmt"
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
		r.pos.y++
	case down:
		r.pos.y--
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

	fmt.Println("*** PART 1 ***")
	program.InitStateFromFile(os.Args[1])

	halted := false

	hull := make(Hull)
	robot := Robot{heading: up}

	for !halted {
		colorUnderRobot := hull[robot.GetPos()].color
		if !halted {
			program.SetReaderFromInts(colorUnderRobot)
			var colorToPaint, directionToTurn int
			colorToPaint, halted = program.RunUntilOutput()
			program.SetReader(nil)
			directionToTurn, halted = program.RunUntilOutput()
			hull.PaintTile(robot.pos, colorToPaint)
			robot.Turn(Direction(directionToTurn))
			robot.MoveForward()
		}
	}
	numPainted := 0
	for _, v := range hull {
		if v.painted {
			numPainted++
		}
	}
	fmt.Println(numPainted)
}
