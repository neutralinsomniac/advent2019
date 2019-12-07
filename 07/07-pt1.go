package main

import (
	"fmt"
	"os"
	"strings"

	//"strings"

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
	cleanAmplifier := intcode.Program{}

	ampA := intcode.Program{}
	ampB := intcode.Program{}
	ampC := intcode.Program{}
	ampD := intcode.Program{}
	ampE := intcode.Program{}

	amps := []*intcode.Program{&ampA, &ampB, &ampC, &ampD, &ampE}

	fmt.Println("*** PART 1 ***")
	cleanAmplifier.InitStateFromFile(os.Args[1])

	phases := []int{0, 1, 2, 3, 4}
	var bestPhase []int
	bestThrust := 0
	for p := make([]int, len(phases)); p[0] < len(p); nextPerm(p) {
		inputSignal := 0
		var thrust int
		ampA.InitStateFromProgram(&cleanAmplifier)
		ampB.InitStateFromProgram(&cleanAmplifier)
		ampC.InitStateFromProgram(&cleanAmplifier)
		ampD.InitStateFromProgram(&cleanAmplifier)
		ampE.InitStateFromProgram(&cleanAmplifier)
		phaseInputs := getPerm(phases, p)
		for i, phase := range phaseInputs {
			ampInput := strings.NewReader(fmt.Sprintf("%d\n%d\n", phase, inputSignal))
			output := amps[i].Run(ampInput)
			inputSignal = output[0]
			thrust = output[0] // on our last loop iteration, this will have the correct value
		}
		// check to see if we reached MAX THRUST
		if thrust > bestThrust {
			copy(bestPhase, phaseInputs)
			bestThrust = thrust
		}
	}
	fmt.Println("best thrust value:", bestThrust)
}