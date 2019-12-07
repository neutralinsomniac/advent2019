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
	amps := []*intcode.Program{
		&intcode.Program{},
		&intcode.Program{},
		&intcode.Program{},
		&intcode.Program{},
		&intcode.Program{},
	}

	fmt.Println("*** PART 1 ***")
	amps[0].InitStateFromFile(os.Args[1])
	// copy ampA to all other amps
	for i, amp := range amps {
		if i != 0 {
			amp.InitStateFromProgram(amps[0])
		}
	}

	phases := []int{0, 1, 2, 3, 4}
	bestThrust := 0
	for p := make([]int, len(phases)); p[0] < len(p); nextPerm(p) {
		inputSignal := 0
		// reset ALL THE AMPS
		for _, amp := range amps {
			amp.Reset()
		}
		phaseInputs := getPerm(phases, p)
		for i, phase := range phaseInputs {
			ampInput := strings.NewReader(fmt.Sprintf("%d\n%d\n", phase, inputSignal))
			output := amps[i].Run(ampInput)
			inputSignal = output[0]
		}
		// check to see if we reached MAX THRUST
		if inputSignal > bestThrust {
			bestThrust = inputSignal
		}
	}
	fmt.Println("best thrust value:", bestThrust)
}
