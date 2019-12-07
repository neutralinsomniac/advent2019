package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Universe map[string]string

func (u Universe) String() string {
	s := ""
	for k, v := range u {
		s += fmt.Sprintf("%s orbits %s\n", k, v)
	}
	return s
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func initStateFromFile(filename string) Universe {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	universe := make(Universe)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		orbits := strings.Split(scanner.Text(), ")")
		// orbits[1] orbits (points to) orbits[0]
		// there should be a unique orbits[1] on every line, so we can always add it to our universe
		universe[orbits[1]] = orbits[0]
	}

	return universe
}

func (u Universe) countOrbits() int {
	numOrbits := 0
	for k := range u {
		for k != "COM" {
			numOrbits++
			k = u[k]
		}
	}
	return numOrbits
}

func main() {
	universe := initStateFromFile("input")
	fmt.Println(universe)

	fmt.Println(universe.countOrbits())

}
