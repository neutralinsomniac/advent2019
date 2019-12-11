package main

import (
	"fmt"
	"os"

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

	// first init from phase inputs
	for i, phase := range phaseInputs {
		amps[i].SetReaderFromInts(phase, inputSignal)
		inputSignal, _ = amps[i].RunUntilOutput()
	}
	// now feedback loop until halt
	i := 0
	for halted := false; halted != true; i++ {
		var tmp int
		// only update our signal if this amp actually returns a signal (instead of halting)
		amps[i%len(amps)].SetReaderFromInts(inputSignal)
		tmp, halted = amps[i%len(amps)].RunUntilOutput()
		if !halted {
			inputSignal = tmp
		}
	}
	result <- inputSignal
}

func main() {
	fmt.Println("*** PART 2 ***")
	baseProg := intcode.Program{}
	baseProg.InitStateFromFile(os.Args[1])
	numWorkers := 0

	results := make(chan int)
	phases := []int{5, 6, 7, 8, 9}
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
