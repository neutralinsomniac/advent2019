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
	chemical ChemicalAmount
	needs    []ChemicalAmount
}

type State struct {
	// chemical name -> Recipe
	book map[string]Recipe
	// chemical name -> amount in supply
	supply           map[string]int
	totalOreProduced int
}

// this takes a *REALLY* long time (~2.5 hours on my vps) BUT it relies 0% on guess-and-check binary searching and just churns on creating fuel for as long as it has the available resources. not an optimal solution, but I'm proud of the approach so I'm keeping it
func (s *State) produce(name string, amount int) {
		if name == "ORE" && s.supply[name] < amount {
			fmt.Println("fuel:", s.supply["FUEL"])
		} else if name == "ORE" && s.supply[name] % 100000 == 0 {
			fmt.Println("ore:", s.supply[name])
		}
		for s.supply[name] < amount {
			for _, need := range s.book[name].needs {
				s.produce(need.name, need.amount)
			}
			s.supply[name] += s.book[name].chemical.amount
		}
		s.supply[name] -= amount
}

func main() {
	fmt.Println("*** PART 2 ***")
	file, err := os.Open(os.Args[1])
	check(err)
	defer file.Close()

	state := State{}
	state.book = make(map[string]Recipe)
	state.supply = make(map[string]int)

	state.supply["ORE"] = 1000000000000
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tmp := strings.Split(line, " => ")
		left := tmp[0]
		right := tmp[1]
		rightChemical := Recipe{}
		fmt.Sscanf(right, "%d %s\n", &rightChemical.chemical.amount, &rightChemical.chemical.name)
		needs := make([]ChemicalAmount, 0)
		for _, text := range strings.Split(left, ", ") {
			need := ChemicalAmount{}
			fmt.Sscanf(text, "%d %s", &need.amount, &need.name)
			need.name = strings.TrimSpace(need.name)
			needs = append(needs, need)
		}
		rightChemical.needs = needs
		state.book[rightChemical.chemical.name] = rightChemical
	}

	state.produce("FUEL", 100000000000)
	fmt.Println(state.totalOreProduced)
}
