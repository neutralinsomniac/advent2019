package main

import (
	"fmt"
	"github.com/neutralinsomniac/advent2019/intcode"
	"os"
	"strings"
)

func main() {
	program := intcode.Program{}

	fmt.Println("*** PART 1 ***")
	program.InitStateFromFile(os.Args[1])
	program.Run(strings.NewReader("1\n"))

	fmt.Println("*** PART 2 ***")
	program.InitStateFromFile(os.Args[1])
	program.Run(strings.NewReader("5\n"))
}