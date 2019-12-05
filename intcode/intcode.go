package intcode

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Opcode int

const (
	Add    Opcode = 1
	Mult          = 2
	Input         = 3
	Output        = 4
	JumpIfTrue    = 5
	JumpIfFalse   = 6
	LessThan      = 7
	Equals        = 8
	Halt          = 99
)

type AddressingMode int

const (
	Ind AddressingMode = 0
	Imm                = 1
)

type Program struct {
	memory []int
	ip     int
}

func (o Opcode) String() string {
	switch o {
	case Add:
		return "Add"
	case Mult:
		return "Mult"
	case Input:
		return "Input"
	case Output:
		return "Output"
	case JumpIfTrue:
		return "JumpIfTrue"
	case JumpIfFalse:
		return "JumpIfFalse"
	case LessThan:
		return "LessThan"
	case Equals:
		return "Equals"
	case Halt:
		return "Halt"
	default:
		return "Unknown"
	}
}

func (a AddressingMode) String() string {
	switch a {
	case Ind:
		return "Ind"
	case Imm:
		return "Imm"
	default:
		return "Unknown"
	}
}

func (p *Program) InitStateFromFile(filename string) {
	dat, err := ioutil.ReadFile(os.Args[1])
	check(err)

	stringArray := strings.Split(string(dat), ",")

	p.memory = make([]int, len(stringArray))
	for i := 0; i < len(stringArray); i++ {
		p.memory[i], err = strconv.Atoi(strings.TrimSpace(stringArray[i]))
		check(err)
	}

	p.ip = 0

	return
}

func (p *Program) SetIp(ip int) {
	p.ip = ip
}

func (p *Program) GetIp() int {
	return p.ip
}

func (p *Program) IncrementIp(amount int) {
	p.ip += amount
}

func (p *Program) GetMemory(index int) int {
	return p.memory[index]
}

func (p *Program) GetOpcode() Opcode {
	return (Opcode)(p.memory[p.ip] % 100)
}

func (p *Program) GetAddressingMode(index int) AddressingMode {
	return (AddressingMode)((p.memory[p.ip] / int(math.Pow10(index+1))) % 10)
}

func (p *Program) GetOutputOperand(index int) *int {
	return &p.memory[p.memory[p.ip+index]]
}

func (p *Program) GetInputOperand(index int) int {
	var inputValue int
	mode := p.GetAddressingMode(index)
	operand := p.memory[p.ip+index]
	switch mode {
	case Ind:
		inputValue = p.memory[operand]
	case Imm:
		inputValue = operand
	default:
		panic(fmt.Sprintf("unknown addressing mode: %v", mode))
	}
	return inputValue
}
