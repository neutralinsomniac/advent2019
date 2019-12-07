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
	text   []int
	memory []int
	ip     int
	halted bool
	reader *bufio.Reader
	output []int
	debug  bool
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

func (p *Program) setDebug(val bool) {
	p.debug = val
}

func (p *Program) Reset() {
	p.ip = 0
	p.output = nil
	p.halted = false
	if len(p.memory) != len(p.text) {
		p.memory = make([]int, len(p.text))
	}
	copy(p.memory, p.text)
}

func (p *Program) InitStateFromFile(filename string) {
	dat, err := ioutil.ReadFile(os.Args[1])
	check(err)

	stringArray := strings.Split(string(dat), ",")

	// copy text section
	p.text = make([]int, len(stringArray))
	for i := 0; i < len(stringArray); i++ {
		p.text[i], err = strconv.Atoi(strings.TrimSpace(stringArray[i]))
		check(err)
	}

	p.Reset()
	return
}

func (p *Program) InitStateFromProgram(other *Program) {
	if len(p.text) != len(other.text) {
		p.text = make([]int, len(other.text))
	}
	copy(p.text, other.text)

	if len(p.memory) != len(other.memory) {
		p.memory = make([]int, len(other.memory))
	}
	copy(p.memory, other.memory)

	if len(p.output) != len(other.output) {
		p.output = make([]int, len(other.output))
	}
	copy(p.output, other.output)

	p.ip = other.ip
	p.halted = other.halted
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

func (p *Program) SetMemory(index int, value int) {
	p.memory[index] = value
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
		if p.debug {
			fmt.Printf("*%d = %d + %d\n", p.GetMemory(p.GetIp()+3), input1, input2)
		}
		*dest = input1 + input2
		p.ip += 4
	case Mult:
		dest := p.GetOutputOperand(3)
		input1 := p.GetInputOperand(1)
		input2 := p.GetInputOperand(2)
		if p.debug {
			fmt.Printf("*%d = %d + %d\n", p.GetMemory(p.GetIp()+3), input1, input2)
		}
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
		if p.debug {
			fmt.Printf("output: %d\n", *dest)
		}
		p.output = append(p.output, *dest)
		p.ip += 2
	case JumpIfTrue:
		input1 := p.GetInputOperand(1)
		input2 := p.GetInputOperand(2)
		if p.debug {
			fmt.Printf("if %d != 0, jmp %d\n", input1, input2)
		}
		if input1 != 0 {
			p.ip = input2
		} else {
			p.ip += 3
		}
	case JumpIfFalse:
		input1 := p.GetInputOperand(1)
		input2 := p.GetInputOperand(2)
		if p.debug {
			fmt.Printf("if %d == 0, jmp %d\n", input1, input2)
		}
		if input1 == 0 {
			p.ip = input2
		} else {
			p.ip += 3
		}
	case LessThan:
		input1 := p.GetInputOperand(1)
		input2 := p.GetInputOperand(2)
		dest := p.GetOutputOperand(3)
		if p.debug {
			fmt.Printf("*%d = (if %d < %d)\n", p.GetMemory(p.GetIp()+3), input1, input2)
		}
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
		if p.debug {
			fmt.Printf("*%d = (if %d == %d)\n", p.GetMemory(p.GetIp()+3), input1, input2)
		}
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

func (p *Program) Run(reader io.Reader) []int {
	p.reader = bufio.NewReader(reader)

	for !p.halted {
		p.Step()
	}

	return p.output
}

/* return output, halted */
func (p *Program) RunUntilOutput(reader io.Reader) (int, bool) {
	p.reader = bufio.NewReader(reader)
	for p.GetOpcode() != Output && !p.halted {
		p.Step()
	}

	if p.halted {
		return 0, p.halted
	}

	// execute the Output opcode
	p.Step()

	// return the last output
	return p.output[len(p.output)-1], p.halted
}
