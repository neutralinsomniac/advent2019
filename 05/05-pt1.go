package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"pintobyte.com/advent2019/intcode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	program := intcode.Program{}
	program.InitStateFromFile(os.Args[1])
	reader := bufio.NewReader(os.Stdin)

	running := true
	for running {

		opcode := program.GetOpcode()
		switch opcode {
		case intcode.Add:
			dest := program.GetOutputOperand(3)
			input1 := program.GetInputOperand(1)
			input2 := program.GetInputOperand(2)
			fmt.Printf("*%d = %d + %d\n", program.GetMemory(program.GetIp()+3), input1, input2)
			*dest = input1 + input2
			program.IncrementIp(4)
		case intcode.Mult:
			dest := program.GetOutputOperand(3)
			input1 := program.GetInputOperand(1)
			input2 := program.GetInputOperand(2)
			fmt.Printf("*%d = %d + %d\n", program.GetMemory(program.GetIp()+3), input1, input2)
			*dest = input1 * input2
			program.IncrementIp(4)
		case intcode.Input:
			dest := program.GetOutputOperand(1)
			valStr, _ := reader.ReadString('\n')
			val, err := strconv.Atoi(valStr[:len(valStr)-1])
			check(err)
			*dest = val
			program.IncrementIp(2)
		case intcode.Output:
			dest := program.GetOutputOperand(1)
			fmt.Printf("output: %d\n", *dest)
			program.IncrementIp(2)
		case intcode.Halt:
			running = false
		default:
			panic(fmt.Sprintf("encountered unknown opcode: %d", opcode))
		}
	}
}
