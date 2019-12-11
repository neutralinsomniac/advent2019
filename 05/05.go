package main

import (
	"fmt"
	"os"

	"github.com/neutralinsomniac/advent2019/intcode"
)

func main() {
	program := intcode.Program{}

	fmt.Println("*** PART 1 ***")
	program.InitStateFromFile(os.Args[1])
	program.SetReaderFromInts(1)
	output := program.Run()
	fmt.Println(output)

	fmt.Println("*** PART 2 ***")
	program.InitStateFromFile(os.Args[1])
	program.SetReaderFromInts(5)
	output = program.Run()
	fmt.Println(output)
}
