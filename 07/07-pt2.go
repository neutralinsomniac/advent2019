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

func main() {
	ampA := intcode.Program{}
	ampB := intcode.Program{}
	ampC := intcode.Program{}
	ampD := intcode.Program{}
	ampE := intcode.Program{}

	amps := []*intcode.Program{&ampA, &ampB, &ampC, &ampD, &ampE}

	fmt.Println("*** PART 2 ***")
	ampA.InitStateFromFile(os.Args[1])
	// init all but ampA
	for _, amp := range amps {
		if amp != &ampA {
			amp.InitStateFromProgram(&ampA)
		}
	}

	phases := []int{5, 6, 7, 8, 9}
	bestThrust := 0
	for p := make([]int, len(phases)); p[0] < len(p); nextPerm(p) {
		inputSignal := 0
		ampA.Reset()
		ampB.Reset()
		ampC.Reset()
		ampD.Reset()
		ampE.Reset()
		phaseInputs := getPerm(phases, p)
		// first init from phase inputs
		for i, phase := range phaseInputs {
			ampInput := strings.NewReader(fmt.Sprintf("%d\n%d\n", phase, inputSignal))
			inputSignal, _ = amps[i].RunUntilOutput(ampInput)
		}
		// now feedback loop until halt
		i := 0
		for halted := false; halted != true; i++ {
			var tmp int
			ampInput := strings.NewReader(fmt.Sprintf("%d\n", inputSignal))
			// only update our signal if this amp actually returns a signal (instead of halting)
			tmp, halted = amps[i%len(amps)].RunUntilOutput(ampInput)
			if !halted {
				inputSignal = tmp
			}
		}
		// check to see if we reached MAX THRUST
		if inputSignal > bestThrust {
			bestThrust = inputSignal
		}
	}
	fmt.Println("best thrust value:", bestThrust)
}
