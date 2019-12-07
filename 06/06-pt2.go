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

func (u Universe) countOrbitalTransfers() int {
	numOrbits := 0
	COMdistanceFromYOU := make(map[string]int)
	// first, walk back home (COM) from YOU
	k := "YOU"
	for k != "COM" {
		COMdistanceFromYOU[k] = numOrbits
		numOrbits++
		k = u[k]
	}

	numOrbits = 0
	// now, walk backwards from SAN until we hit a path that YOU also hit
	k = "SAN"
	for COMdistanceFromYOU[k] == 0 {
		numOrbits++
		k = u[k]
	}

	return numOrbits + COMdistanceFromYOU[k] - 2 // subtract 2 for the extra path traversals on SAN and YOU
}

func main() {
	universe := initStateFromFile("input")

	fmt.Println(universe.countOrbitalTransfers())

}
