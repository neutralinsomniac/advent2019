package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/neutralinsomniac/advent2019/intcode"
)

func nextPerm(p []int) {
	for i := len(p) - 1; i >= 0; i-- {
		if i == 0 || p[i] < len(p)-i-1 {
			p[i]++
			return
		}
		p[i] = 0
	}
}

func getPerm(orig, p []int) []int {
	result := append([]int{}, orig...)
	for i, v := range p {
		result[i], result[i+v] = result[i+v], result[i]
	}
	return result
}

func work(baseProg *intcode.Program, phaseInputs []int, result chan<- int) {
	amps := []*intcode.Program{
		&intcode.Program{},
		&intcode.Program{},
		&intcode.Program{},
		&intcode.Program{},
		&intcode.Program{},
	}

	// copy baseProg to all other amps
	for _, amp := range amps {
		amp.InitStateFromProgram(baseProg)
	}

	inputSignal := 0
	for i, phase := range phaseInputs {
		ampInput := strings.NewReader(fmt.Sprintf("%d\n%d\n", phase, inputSignal))
		output := amps[i].Run(ampInput)
		inputSignal = output[0]
	}
	result <- inputSignal
}

func main() {
	fmt.Println("*** PART 1 ***")
	baseProg := intcode.Program{}
	baseProg.InitStateFromFile(os.Args[1])
	numWorkers := 0

	results := make(chan int)
	phases := []int{0, 1, 2, 3, 4}
	bestThrust := 0
	for p := make([]int, len(phases)); p[0] < len(p); nextPerm(p) {
		phaseInputs := getPerm(phases, p)
		go work(&baseProg, phaseInputs, results)
		numWorkers++
	}
	for i := 0; i < numWorkers; i++ {
		result := <-results
		// check to see if we reached MAX THRUST
		if result > bestThrust {
			bestThrust = result
		}
	}
	fmt.Println("best thrust value:", bestThrust)
}
