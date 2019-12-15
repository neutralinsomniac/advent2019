package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type ChemicalAmount struct {
	name   string
	amount int
}

func (c ChemicalAmount) String() string {
	return fmt.Sprintf("name: %s amount: %d", c.name, c.amount)
}

type Recipe struct {
	produces ChemicalAmount
	needs    []ChemicalAmount
}

type State struct {
	// chemical name -> Recipe
	book map[string]Recipe
	// chemical name -> amount in supply
	supply map[string]int
}

func (s *State) produce(name string, amount int) {
	for s.supply[name] < amount {
		// produce this shit
		s.supply[name] += s.book[name].produces.amount
		for _, need := range s.book[name].needs {
			s.supply[need.name] -= need.amount
		}
	}
	for _, need := range s.book[name].needs {
		if s.supply[need.name] < 0 {
			if need.name != "ORE" {
				s.produce(need.name, 0)
			}
		}
	}
}

func main() {
	fmt.Println("*** PART 1 ***")
	file, err := os.Open(os.Args[1])
	check(err)
	defer file.Close()

	state := State{}
	state.book = make(map[string]Recipe)
	state.supply = make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tmp := strings.Split(line, " => ")
		left := tmp[0]
		right := tmp[1]
		rightChemical := Recipe{}
		fmt.Sscanf(right, "%d %s\n", &rightChemical.produces.amount, &rightChemical.produces.name)
		needs := make([]ChemicalAmount, 0)
		for _, text := range strings.Split(left, ", ") {
			need := ChemicalAmount{}
			fmt.Sscanf(text, "%d %s", &need.amount, &need.name)
			need.name = strings.TrimSpace(need.name)
			needs = append(needs, need)
		}
		rightChemical.needs = needs
		state.book[rightChemical.produces.name] = rightChemical
	}

	state.produce("FUEL", 1)
	fmt.Println(-state.supply["ORE"])
}
