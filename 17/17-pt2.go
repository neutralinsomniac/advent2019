package main

import (
	"fmt"
	"os"

	"github.com/neutralinsomniac/advent2019/intcode"
)

func main() {
	program := intcode.Program{}

	program.InitStateFromFile(os.Args[1])
	program.SetMemory(0, 2)

	fmt.Println("*** PART 2 ***")

	instructions := "A,B,A,C,A,B,C,A,B,C\nR,12,R,4,R,10,R,12\nR,6,L,8,R,10\nL,8,R,4,R,4,R,6\nn\n"
	ints := make([]int, len(instructions))
	for i, c := range instructions {
		ints[i] = int(c)
	}
	program.SetReaderFromInts(ints...)
	output := program.Run()

	for _, c := range output[:len(output)-1] {
		fmt.Printf("%c", byte(c))
	}
	fmt.Println("dust:", output[len(output)-1])
}
