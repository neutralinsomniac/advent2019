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

type Status int

const (
	wall  Status = 0
	moved        = 1
	won          = 2
)

type Tile struct {
	status  Status
	visited bool
}

type World map[Coord]Tile

func (world World) Print() {
	// first figure out our world boundaries
	minX := math.MaxInt32
	minY := math.MaxInt32
	maxX := -1 * math.MaxInt32
	maxY := -1 * math.MaxInt32
	for coord, _ := range world {
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
	fmt.Printf("\033[H")
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if world[Coord{x, y}].visited == false {
				fmt.Printf(" ")
			} else {
				switch world[Coord{x, y}].status {
				case wall:
					fmt.Printf("#")
				case moved:
					fmt.Printf(".")
				case won:
					fmt.Printf("!")
				}
			}
		}
		fmt.Printf("\n")
	}
}

type Direction int

const (
	north Direction = iota + 1
	south
	west
	east
)

type Robot struct {
	pos Coord
}

func (r Robot) String() string {
	return "@"
}

func Find(slice []Coord, coord Coord) bool {
	for _, item := range slice {
		if item == coord {
			return true
		}
	}
	return false
}

func Explore(program intcode.Program, world World, curLocation Coord, crumbs []Coord) {
	newProg := intcode.Program{}

	for _, direction := range []Direction{north, east, south, west} {
		//world.Print()
		var exploreLoc Coord
		switch direction {
		case north:
			exploreLoc = Coord{curLocation.x, curLocation.y - 1}
		case east:
			exploreLoc = Coord{curLocation.x + 1, curLocation.y}
		case south:
			exploreLoc = Coord{curLocation.x, curLocation.y + 1}
		case west:
			exploreLoc = Coord{curLocation.x - 1, curLocation.y}
		}
		// no backtracking
		if Find(crumbs, exploreLoc) == false {
			newProg.InitStateFromProgram(&program)
			newProg.SetReaderFromInts(int(direction))
			status, _ := newProg.RunUntilOutput()
			world[exploreLoc] = Tile{visited: true, status: Status(status)}
			switch Status(status) {
			case wall:
				continue
			case moved:
				crumbs = append(crumbs, curLocation)
				Explore(newProg, world, exploreLoc, crumbs)
			case won:
				fmt.Println("FOUND OXYGEN AFTER ", len(crumbs), "STEPS")
				return
			}
		}
	}
}

func main() {
	program := intcode.Program{}

	fmt.Println("*** PART 1 ***")
	program.InitStateFromFile(os.Args[1])

	world := make(World)
	world[Coord{}] = Tile{visited: true, status: moved}

	//fmt.Printf("\033[2J;\033[H")

	Explore(program, world, Coord{}, make([]Coord, 0))
}
