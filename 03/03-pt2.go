package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	Wire1 uint8 = 1 << iota
	Wire2
)

type direction int

const (
	up direction = iota
	down
	right
	left
)

type movement struct {
	direction direction
	distance  int
}

type point struct {
	x, y int
}

type footprint struct {
	wireId uint8
	distance int
}

func (p point) String() string {
	return fmt.Sprintf("x: %v, y: %v", p.x, p.y)
}

func (m movement) String() string {
	var dir string
	switch m.direction {
	case up:
		dir = "up"
	case down:
		dir = "down"
	case left:
		dir = "left"
	case right:
		dir = "right"
	}

	return fmt.Sprintf("direction: %v, distance: %v", dir, m.distance)
}

func createMovementsFromStrings(movementString string) []movement {
	movementStringArray := strings.Split(movementString, ",")
	movements := make([]movement, len(movementStringArray))

	var err error

	for i := range movementStringArray {
		switch movementStringArray[i][0] {
		case 'U':
			movements[i].direction = up
		case 'D':
			movements[i].direction = down
		case 'L':
			movements[i].direction = left
		case 'R':
			movements[i].direction = right
		}
		movements[i].distance, err = strconv.Atoi(movementStringArray[i][1:])
		check(err)
	}

	return movements
}

func initStateFromFile(filename string) ([]movement, []movement) {
	dat, err := ioutil.ReadFile(os.Args[1])
	check(err)

	bothWires := strings.Split(string(dat), "\n")

	wire1 := createMovementsFromStrings(bothWires[0])
	wire2 := createMovementsFromStrings(bothWires[1])

	return wire1, wire2
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fuckingMOVE(world map[point]footprint, start *point, distance *int, movement movement, wireId uint8, isects *[]int) {
	fmt.Println(start, movement)

	switch movement.direction {
	case up:
		for i := 0; i < movement.distance; i++ {
			start.y += 1
			*distance += 1
			if f, ok := world[*start]; ok {
				if f.wireId | wireId != f.wireId {
					f.distance += *distance
					f.wireId |= wireId
					world[*start] = f
					*isects = append(*isects, f.distance)
				}
			} else {
				f := footprint{distance: *distance, wireId: wireId}
				world[*start] = f
			}
		}
	case down:
		for i := 0; i < movement.distance; i++ {
			start.y -= 1
			*distance += 1
			if f, ok := world[*start]; ok {
				if f.wireId | wireId != f.wireId {
					f.distance += *distance
					f.wireId |= wireId
					world[*start] = f
				}
			} else {
				f := footprint{distance: *distance, wireId: wireId}
				world[*start] = f
			}

		}
	case right:
		for i := 0; i < movement.distance; i++ {
			start.x += 1
			*distance += 1
			if f, ok := world[*start]; ok {
				if f.wireId | wireId != f.wireId {
					f.distance += *distance
					f.wireId |= wireId
					world[*start] = f
				}
			} else {
				f := footprint{distance: *distance, wireId: wireId}
				world[*start] = f
			}

		}
	case left:
		for i := 0; i < movement.distance; i++ {
			start.x -= 1
			*distance += 1
			if f, ok := world[*start]; ok {
				if f.wireId | wireId != f.wireId {
					f.distance += *distance
					f.wireId |= wireId
					world[*start] = f
				}
			} else {
				f := footprint{distance: *distance, wireId: wireId}
				world[*start] = f
			}

		}
	}
}

func findIntersections(world map[point]uint8) []point {
	var points []point

	for point, wireId := range world {
		if wireId == Wire1 | Wire2 {
			points = append(points, point)
		}
	}

	return points
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattan(start point, end point) int {
	return abs(start.x - end.x) + abs(start.y - end.y)
}

func findShortestManhattanDistance(start point, points []point) int {
	var closestDistance = math.MaxInt32

	if len(points) == 0 {
		return 0
	}

	for _, point := range points {
		dist := manhattan(start, point)
		if dist < closestDistance {
			closestDistance = dist
		}
	}
	return closestDistance
}

func main() {
	wire1, wire2 := initStateFromFile(os.Args[1])

	theWholeWorld := make(map[point]footprint)

	startingPoint := point{}
	distance := 0

	var isects []int

	for _, movement := range wire1 {
		fuckingMOVE(theWholeWorld, &startingPoint, &distance, movement, Wire1, &isects)
	}

	startingPoint = point{}
	distance = 0

	for _, movement := range wire2 {
		fuckingMOVE(theWholeWorld, &startingPoint, &distance, movement, Wire2, &isects)
	}

	fmt.Println(isects)
}
