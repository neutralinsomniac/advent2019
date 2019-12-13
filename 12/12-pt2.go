package main

import (
	"bufio"
	"fmt"
	"os"
)

type vector struct {
	x, y, z int
}

type planet struct {
	pos      vector
	vel      vector
	startpos vector
	startvel vector
}

func (p planet) String() string {
	return fmt.Sprintf("pos: %v, vel: %v", p.pos, p.vel)
}

type universe struct {
	planets                []planet
	minX, minY, maxX, maxY int
}

func (u *universe) Print() {
	for _, p := range u.planets {
		if p.pos.x > u.maxX {
			u.maxX = p.pos.x
		}
		if p.pos.x < u.minX {
			u.minX = p.pos.x
		}
		if p.pos.y > u.maxY {
			u.maxY = p.pos.y
		}
		if p.pos.y < u.minY {
			u.minY = p.pos.y
		}
	}
	for y := u.minY; y <= u.maxY; y++ {
		for x := u.minX; x <= u.maxX; x++ {
			for i, p := range u.planets {
				if p.pos.x == x && p.pos.y == y {
					fmt.Printf("%d", i)
				} else {
					fmt.Printf(" ")
				}
			}
		}
		fmt.Printf("\n")
	}
}

func (u universe) applyGravityToPlanet(index int) {
	for j := range u.planets {
		if j != index {
			u.planets[index].applyGravity(u.planets[j])
		}
	}
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
		planet.startpos = planet.pos
		planet.startvel = planet.vel
		universe.planets = append(universe.planets, planet)
	}

	return universe
}

// greatest common divisor (GCD) via Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}
	return result
}

func main() {
	fmt.Println("*** PART 2 ***")

	universe := InitStateFromFile(os.Args[1])

	num_steps := 0
	x_period := 0
	y_period := 0
	z_period := 0
	last_x := 0
	last_y := 0
	last_z := 0

	//fmt.Printf("\033[2J;\033[H")

	for x_period == 0 || y_period == 0 || z_period == 0 {
		//fmt.Printf("\033[H")

		//universe.Print()
		//time.Sleep(100 * time.Millisecond)

		for j := range universe.planets {
			universe.applyGravityToPlanet(j)
		}
		for j := range universe.planets {
			universe.planets[j].applyVelocity()
		}

		xs_aligned := true
		for j := range universe.planets {
			if universe.planets[j].vel.x != universe.planets[j].startvel.x || universe.planets[j].pos.x != universe.planets[j].startpos.x {
				xs_aligned = false
				break
			}
		}

		if xs_aligned {
			if last_x != 0 {
				x_period = num_steps - last_x
			}
			last_x = num_steps
		}

		ys_aligned := true
		for j := range universe.planets {
			if universe.planets[j].vel.y != universe.planets[j].startvel.y || universe.planets[j].pos.y != universe.planets[j].startpos.y {
				ys_aligned = false
				break
			}
		}

		if ys_aligned {
			if last_y != 0 {
				y_period = num_steps - last_y
			}
			last_y = num_steps
		}

		zs_aligned := true
		for j := range universe.planets {
			if universe.planets[j].vel.z != universe.planets[j].startvel.z || universe.planets[j].pos.z != universe.planets[j].startpos.z {
				zs_aligned = false
				break
			}
		}

		if zs_aligned {
			if last_z != 0 {
				z_period = num_steps - last_z
			}
			last_z = num_steps
		}

		num_steps++
	}

	fmt.Println(lcm(x_period, y_period, z_period))
}
