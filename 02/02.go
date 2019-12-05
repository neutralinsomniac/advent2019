package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile(os.Args[1])
	check(err)

	stringArray := strings.Split(string(dat), ",")

	fmt.Println(stringArray)
	state := make([]int, len(stringArray))
	for i := 0; i < len(stringArray); i++ {
		state[i], err = strconv.Atoi(strings.TrimSpace(stringArray[i]))
		check(err)
	}
	fmt.Println(state)

	fmt.Println(len(state))
	for i := 0; i < len(state); i += 4 {
		opcode := state[i]
		switch opcode {
		case 1:
			state[state[i+3]] = state[state[i+1]] + state[state[i+2]]
		case 2:
			state[state[i+3]] = state[state[i+1]] * state[state[i+2]]
		case 99:
			break
		}
	}
	fmt.Println(state)
}
