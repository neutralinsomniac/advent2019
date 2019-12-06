package intcode

import (
	"bufio"
	"fmt"
	"io"
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
	Add         Opcode = 1
	Mult               = 2
	Input              = 3
	Output             = 4
	JumpIfTrue         = 5
	JumpIfFalse        = 6
	LessThan           = 7
	Equals             = 8
	Halt               = 99
)

type AddressingMode int

const (
	Ind AddressingMode = 0
	Imm                = 1
)

type Program struct {
	memory []int
	ip     int
	halted bool
	reader *bufio.Reader
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
	return Opcode(p.memory[p.ip] % 100)
}

func (p *Program) GetAddressingMode(index int) AddressingMode {
	return AddressingMode((p.memory[p.ip] / int(math.Pow10(index+1))) % 10)
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

func (p *Program) Step() {
	opcode := p.GetOpcode()
	switch opcode {
	case Add:
		dest := p.GetOutputOperand(3)
		input1 := p.GetInputOperand(1)
		input2 := p.GetInputOperand(2)
		fmt.Printf("*%d = %d + %d\n", p.GetMemory(p.GetIp()+3), input1, input2)
		*dest = input1 + input2
		p.ip += 4
	case Mult:
		dest := p.GetOutputOperand(3)
		input1 := p.GetInputOperand(1)
		input2 := p.GetInputOperand(2)
		fmt.Printf("*%d = %d + %d\n", p.GetMemory(p.GetIp()+3), input1, input2)
		*dest = input1 * input2
		p.ip += 4
	case Input:
		dest := p.GetOutputOperand(1)
		valStr, _ := p.reader.ReadString('\n')
		val, err := strconv.Atoi(valStr[:len(valStr)-1])
		check(err)
		*dest = val
		p.ip += 2
	case Output:
		dest := p.GetOutputOperand(1)
		fmt.Printf("output: %d\n", *dest)
		p.ip += 2
	case JumpIfTrue:
		input1 := p.GetInputOperand(1)
		input2 := p.GetInputOperand(2)
		fmt.Printf("if %d != 0, jmp %d\n", input1, input2)
		if input1 != 0 {
			p.ip = input2
		} else {
			p.ip += 3
		}
	case JumpIfFalse:
		input1 := p.GetInputOperand(1)
		input2 := p.GetInputOperand(2)
		fmt.Printf("if %d == 0, jmp %d\n", input1, input2)
		if input1 == 0 {
			p.ip = input2
		} else {
			p.ip += 3
		}
	case LessThan:
		input1 := p.GetInputOperand(1)
		input2 := p.GetInputOperand(2)
		dest := p.GetOutputOperand(3)
		fmt.Printf("*%d = (if %d < %d)\n", p.GetMemory(p.GetIp()+3), input1, input2)
		if input1 < input2 {
			*dest = 1
		} else {
			*dest = 0
		}
		p.ip += 4
	case Equals:
		input1 := p.GetInputOperand(1)
		input2 := p.GetInputOperand(2)
		dest := p.GetOutputOperand(3)
		fmt.Printf("*%d = (if %d == %d)\n", p.GetMemory(p.GetIp()+3), input1, input2)
		if input1 == input2 {
			*dest = 1
		} else {
			*dest = 0
		}
		p.ip += 4
	case Halt:
		p.halted = true
	default:
		panic(fmt.Sprintf("encountered unknown opcode: %d", opcode))
	}
}

func (p *Program) StepBy(steps int) {
	for i := 0; i < steps; i++ {
		p.Step()
	}
}

func (p *Program) Run(reader io.Reader) {
	p.halted = false
	p.ip = 0
	p.reader = bufio.NewReader(reader)

	for !p.halted {
		p.Step()
	}
}
