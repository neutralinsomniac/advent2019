package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Coord struct {
	x int
	y int
}

// lol globals
var universeWidth int
var universeHeight int

type Universe map[Coord]bool

func (u Universe) String() string {
	var sb strings.Builder
	for y := 0; y < universeHeight; y++ {
		for x := 0; x < universeWidth; x++ {
			if u[Coord{x, y}] {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func InitStateFromFile(filename string) Universe {
	dat, err := ioutil.ReadFile(os.Args[1])
	check(err)

	universe := make(Universe)
	var x, y int
	for _, c := range dat {
		switch c {
		case '#':
			universe[Coord{x: x, y: y}] = true
		case '\n':
			y++
			universeWidth = x
			x = 0
			continue
		}
		x++
	}
	universeHeight = y
	return universe
}

func calcDistance(start, end Coord) float64 {
	return math.Sqrt(math.Pow(float64(end.x-start.x), 2) + math.Pow(float64(end.y-start.y), 2))
}

func calcAngle(start, end Coord) float64 {
	diffx := float64(end.x - start.x)
	diffy := float64(start.y - end.y)
	degrees := math.Atan2(diffy, diffx) * (180 / math.Pi)
	if degrees < 0 {
		degrees += 360
	}
	// rotate that shit
	degrees = degrees - 90
	if degrees < 0 {
		degrees += 360
	}
	// and walk clockwise instead of anti-clockwise
	if degrees != 0 {
		degrees = 360 - degrees
	}
	return degrees
}

func main() {
	fmt.Println("*** PART 2 ***")

	universe := InitStateFromFile("input")

	// first find the best monitor
	bestMonitor := 0
	// angle, distance, Coord
	var bestShitToVaporize map[float64]map[float64]Coord
	var shitToVaporize map[float64]map[float64]Coord
	for start, _ := range universe {
		shitToVaporize = make(map[float64]map[float64]Coord)
		for end, _ := range universe {
			if start == end {
				continue
			}
			angle := calcAngle(start, end)
			distance := calcDistance(start, end)
			if shitToVaporize[angle] == nil {
				shitToVaporize[angle] = make(map[float64]Coord)
			}

			shitToVaporize[angle][distance] = end
		}
		if len(shitToVaporize) > bestMonitor {
			bestMonitor = len(shitToVaporize)
			bestShitToVaporize = shitToVaporize
		}
	}

	sortedAngles := make([]float64, 0, len(bestShitToVaporize))
	for k := range bestShitToVaporize {
		sortedAngles = append(sortedAngles, k)
	}

	sort.Float64s(sortedAngles)
	fmt.Printf("\033[2J;\033[H")

	numVaporized := 0
	for vapedSomething := true; vapedSomething; {
		vapedSomething = false
		for _, angle := range sortedAngles {
			if len(bestShitToVaporize[angle]) > 0 {
				numVaporized++
				sortedDistances := make([]float64, 0, len(bestShitToVaporize[angle]))
				for distance := range bestShitToVaporize[angle] {
					sortedDistances = append(sortedDistances, distance)
				}
				sort.Float64s(sortedDistances)
				coord := bestShitToVaporize[angle][sortedDistances[0]]

				if numVaporized == 200 {
					fmt.Println("WINNER", coord.x*100+coord.y)
				}
				delete(bestShitToVaporize[angle], sortedDistances[0])
				delete(universe, coord)
				fmt.Printf("\033[H")
				fmt.Println(universe)
				time.Sleep(50 * time.Millisecond)
				vapedSomething = true
			}
		}
	}
}
