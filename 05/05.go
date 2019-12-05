package main

import (
	"os"
	"github.com/neutralinsomniac/advent2019/intcode"
)

func main() {
	program := intcode.Program{}
	program.InitStateFromFile(os.Args[1])
	program.Run()
}
