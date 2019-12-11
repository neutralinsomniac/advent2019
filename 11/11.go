package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/neutralinsomniac/advent2019/intcode"
)

type Coord struct {
	x, y int
}

type Hull map[Coord]int

type Heading int

const (
	up Heading = iota
	down
	left
	right
)

type Direction int

const (
	turnLeft Direction = 0
	turnRight = 1
)

const (
	black = 0
	white = 1
)

type Robot struct {
	heading Heading
	pos Coord
}

func (r *Robot) Turn(
func main() {
	program := intcode.Program{}

	fmt.Println("*** PART 1 ***")
	program.InitStateFromFile(os.Args[1])

	halted := false

	hull := make(Hull)
	robot := Robot{heading: up}

	for !halted {
		
	}

	output := program.Run(strings.NewReader("1\n"))
	fmt.Println(output)

}
