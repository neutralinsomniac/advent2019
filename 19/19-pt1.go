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

type World map[Coord]bool

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
			if world[Coord{x, y}] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	fmt.Println("*** PART 1 ***")

	program := intcode.Program{}
	program.InitStateFromFile(os.Args[1])

	world := make(World)

	sum := 0
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			program.Reset()
			program.SetReaderFromInts(x, y)
			out, _ := program.RunUntilOutput()
			if out == 1 {
				world[Coord{x, y}] = true
				sum++
			}
		}
	}
	fmt.Println(sum)
}
