package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/neutralinsomniac/advent2019/intcode"
)

func main() {
	program := intcode.Program{}

	fmt.Println("*** PART 1 ***")
	program.InitStateFromFile(os.Args[1])
	output := program.Run(strings.NewReader("1\n"))
	fmt.Println(output)

	fmt.Println("*** PART 2 ***")
	program.Reset()
	output = program.Run(strings.NewReader("2\n"))
	fmt.Println(output)
}
