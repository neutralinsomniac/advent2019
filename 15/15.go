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
	wall       Status = 0
	moved             = 1
	won               = 2
	oxygenated        = 3
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
				case oxygenated:
					fmt.Printf("O")
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
				fmt.Println("FOUND OXYGEN AFTER", len(crumbs), "STEPS")
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

	fmt.Println("*** PART 2 ***")
	var start Coord
	for coord, tile := range world {
		if tile.status == won {
			start = coord
			break
		}
	}

	toFill := make(map[Coord]bool)
	toSearch := make(map[Coord]bool)

	toSearch[start] = true
	world[start] = Tile{status: oxygenated}

	minutes := 0
	for len(toSearch) > 0 {
		minutes++
		toFill = make(map[Coord]bool)
		for coord := range toSearch {
			for _, coord := range []Coord{Coord{coord.x, coord.y - 1}, Coord{coord.x, coord.y + 1}, Coord{coord.x - 1, coord.y}, Coord{coord.x + 1, coord.y}} {
				if world[coord].status == moved {
					toFill[coord] = true
				}
			}

		}
		for coord := range toFill {
			world[coord] = Tile{visited: true, status: oxygenated}
		}
		toSearch = toFill
		//world.Print()
	}
	// subtract 1 from minutes because our last run just verifies that the whole ship is filled with oxygen
	fmt.Println("oxygen fill took", minutes-1, "minutes")
}
