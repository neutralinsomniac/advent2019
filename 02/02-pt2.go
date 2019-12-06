package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func initStateFromFile(filename string) []int {
	dat, err := ioutil.ReadFile(os.Args[1])
	check(err)

	stringArray := strings.Split(string(dat), ",")

	state := make([]int, len(stringArray))
	for i := 0; i < len(stringArray); i++ {
		state[i], err = strconv.Atoi(strings.TrimSpace(stringArray[i]))
		check(err)
	}
	return state
}

func main() {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			state := initStateFromFile(os.Args[1])
			running := true
			state[1] = noun
			state[2] = verb
			i := 0
			for running {
				opcode := state[i]
				switch opcode {
				case 1:
					state[state[i+3]] = state[state[i+1]] + state[state[i+2]]
					i += 4
				case 2:
					state[state[i+3]] = state[state[i+1]] * state[state[i+2]]
					i += 4
				case 99:
					running = false
				}
			}
			if state[0] == 19690720 {
				fmt.Println(state)
				return
			}
		}
	}
}
