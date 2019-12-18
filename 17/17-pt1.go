package main

import (
	"fmt"
	"os"

	"github.com/neutralinsomniac/advent2019/intcode"
)

type Coord struct {
	x, y int
}

type Material byte

const (
	scaffolding Material = '#'
	space                = '.'
)

type World map[Coord]Material

func main() {
	fmt.Println("*** PART 1 ***")

	program := intcode.Program{}

	program.InitStateFromFile(os.Args[1])

	world := make(World)

	var coord Coord

	var halted bool
	for !halted {
		var material int
		material, halted = program.RunUntilOutput()
		switch byte(material) {
		case byte(scaffolding), byte(space), '<', '>', '^', 'v':
			world[coord] = Material(material)
			coord.x++
		case '\n':
			coord.x = 0
			coord.y++
		}
	}

	sum := 0
	for coord, material := range world {
		if material == scaffolding {
			if world[Coord{coord.x + 1, coord.y}] == scaffolding &&
				world[Coord{coord.x - 1, coord.y}] == scaffolding &&
				world[Coord{coord.x, coord.y + 1}] == scaffolding &&
				world[Coord{coord.x, coord.y - 1}] == scaffolding {
				sum += coord.x * coord.y
			}
		}
	}
	fmt.Println(sum)
}
