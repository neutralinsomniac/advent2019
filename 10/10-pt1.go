package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Coord struct {
	x, y int
}

func InitStateFromFile(filename string) map[Coord]bool {
	dat, err := ioutil.ReadFile(filename)
	check(err)

	universe := make(map[Coord]bool)
	var x, y int
	for _, c := range dat {
		switch c {
		case '#':
			universe[Coord{x: x, y: y}] = true
		case '\n':
			x++
			y = 0
		}
		y++
	}
	return universe
}

func calcAngle(start, end Coord) float64 {
	diffx := float64(end.x - start.x)
	diffy := float64(end.y - start.y)
	return math.Atan2(diffy, diffx) * (180 / math.Pi)
}

func main() {
	fmt.Println("*** PART 1 ***")

	universe := InitStateFromFile("input")

	bestMonitor := 0
	for start, _ := range universe {
		angles := make(map[float64]bool)
		for end, _ := range universe {
			if start != end {
				angles[calcAngle(start, end)] = true
			}
		}
		if len(angles) > bestMonitor {
			bestMonitor = len(angles)
		}
	}
	fmt.Println(bestMonitor)
}
