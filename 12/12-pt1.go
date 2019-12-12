package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type vector struct {
	x, y, z int
}

type planet struct {
	pos vector
	vel vector
}

func (p planet) String() string {
	return fmt.Sprintf("pos: %v, vel: %v", p.pos, p.vel)
}

type universe struct {
	planets []planet
}

func (u universe) String() string {
	var sb strings.Builder

	for _, p := range u.planets {
		fmt.Fprintf(&sb, "%s\n", p)
	}
	return sb.String()
}

func (u universe) applyGravityToPlanet(index int) {
	for j := range u.planets {
		if j != index {
			u.planets[index].applyGravity(u.planets[j])
		}
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (p planet) calculateKineticEnergy() int {
	return abs(p.vel.x) + abs(p.vel.y) + abs(p.vel.z)
}

func (p planet) calculatePotentialEnergy() int {
	return abs(p.pos.x) + abs(p.pos.y) + abs(p.pos.z)
}

func (p planet) calculateTotalEnergy() int {
	return p.calculateKineticEnergy() * p.calculatePotentialEnergy()
}

func (u universe) calculateTotalEnergy() int {
	sum := 0
	for _, p := range u.planets {
		sum += p.calculateTotalEnergy()
	}
	return sum
}

func (p *planet) applyGravity(other planet) {
	if p.pos.x < other.pos.x {
		p.vel.x++
	} else if p.pos.x > other.pos.x {
		p.vel.x--
	}
	if p.pos.y < other.pos.y {
		p.vel.y++
	} else if p.pos.y > other.pos.y {
		p.vel.y--
	}
	if p.pos.z < other.pos.z {
		p.vel.z++
	} else if p.pos.z > other.pos.z {
		p.vel.z--
	}
}

func (p *planet) applyVelocity() {
	p.pos.x += p.vel.x
	p.pos.y += p.vel.y
	p.pos.z += p.vel.z
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func InitStateFromFile(filename string) universe {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	var universe universe

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		planet := planet{}
		fmt.Sscanf(scanner.Text(), "<x=%d, y=%d, z=%d>\n", &planet.pos.x, &planet.pos.y, &planet.pos.z)
		universe.planets = append(universe.planets, planet)
	}

	return universe
}

func main() {
	fmt.Println("*** PART 1 ***")

	universe := InitStateFromFile(os.Args[1])

	numSteps := 1000

	for i := 0; i < numSteps; i++ {
		for j := range universe.planets {
			universe.applyGravityToPlanet(j)
		}
		for j := range universe.planets {
			universe.planets[j].applyVelocity()
		}
	}

	fmt.Println(universe.calculateTotalEnergy())
}
