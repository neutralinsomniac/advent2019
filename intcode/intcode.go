package intcode

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"math"
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
	AdjustBP           = 9
	Halt               = 99
)

type AddressingMode int

const (
	Ind AddressingMode = 0
	Imm                = 1
	Rel                = 2
)

type Program struct {
	text   []int
	memory map[int]int
	ip     int
	bp     int
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
	case AdjustBP:
		return "AdjustBP"
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
	case Rel:
		return "Rel"
	default:
		return "Unknown"
	}
}

func (p *Program) SetDebug(val bool) {
	p.debug = val
}

func (p *Program) SetReader(reader io.Reader) {
	p.reader = bufio.NewReader(reader)
}

func (p *Program) SetReaderFromInts(ints ...int) {
	var sb strings.Builder
	for _, i := range ints {
		fmt.Fprintf(&sb, "%d\n", i)
	}
	p.SetReader(strings.NewReader(sb.String()))
}

func (p *Program) Reset() {
	p.ip = 0
	p.bp = 0
	p.output = nil
	p.reader = nil
	p.halted = false
	p.memory = make(map[int]int)
	for i := range p.text {
		p.memory[i] = p.text[i]
	}
}

func (p *Program) InitStateFromFile(filename string) {
	dat, err := ioutil.ReadFile(filename)
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

	p.memory = make(map[int]int)
	for k, v := range other.memory {
		p.memory[k] = v
	}

	if len(p.output) != len(other.output) {
		p.output = make([]int, len(other.output))
	}
	copy(p.output, other.output)

	p.ip = other.ip
	p.bp = other.bp
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

func (p *Program) GetOutputOperand(index int) int {
	mode := p.GetAddressingMode(index)
	operand := p.memory[p.ip+index]
	switch mode {
	case Ind:
		return operand
	case Imm:
		panic(fmt.Sprintf("tried to use immediate addressing mode for output param"))
	case Rel:
		return p.bp + operand
	default:
		panic(fmt.Sprintf("unknown addressing mode: %v", mode))
	}
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
	case Rel:
		inputValue = p.memory[p.bp+operand]
	default:
		panic(fmt.Sprintf("unknown addressing mode: %v", mode))
	}
	return inputValue
}

func (p *Program) Step() {
	opcode := p.GetOpcode()
	switch opcode {
	case Add:
		destIndex := p.GetOutputOperand(3)
		input1 := p.GetInputOperand(1)
		input2 := p.GetInputOperand(2)
		if p.debug {
			fmt.Printf("*%d = %d + %d\n", p.GetMemory(p.GetIp()+3), input1, input2)
		}
		p.memory[destIndex] = input1 + input2
		p.ip += 4
	case Mult:
		destIndex := p.GetOutputOperand(3)
		input1 := p.GetInputOperand(1)
		input2 := p.GetInputOperand(2)
		if p.debug {
			fmt.Printf("*%d = %d + %d\n", p.GetMemory(p.GetIp()+3), input1, input2)
		}
		p.memory[destIndex] = input1 * input2
		p.ip += 4
	case Input:
		if p.reader == nil {
			panic("encountered input instruction and no reader is set")
		}
		destIndex := p.GetOutputOperand(1)
		valStr, _ := p.reader.ReadString('\n')
		val, err := strconv.Atoi(valStr[:len(valStr)-1])
		check(err)
		p.memory[destIndex] = val
		p.ip += 2
	case Output:
		src := p.GetInputOperand(1)
		if p.debug {
			fmt.Printf("output: %d\n", src)
		}
		p.output = append(p.output, src)
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
		destIndex := p.GetOutputOperand(3)
		if p.debug {
			fmt.Printf("*%d = (if %d < %d)\n", p.GetMemory(p.GetIp()+3), input1, input2)
		}
		if input1 < input2 {
			p.memory[destIndex] = 1
		} else {
			p.memory[destIndex] = 0
		}
		p.ip += 4
	case Equals:
		input1 := p.GetInputOperand(1)
		input2 := p.GetInputOperand(2)
		destIndex := p.GetOutputOperand(3)
		if p.debug {
			fmt.Printf("*%d = (if %d == %d)\n", p.GetMemory(p.GetIp()+3), input1, input2)
		}
		if input1 == input2 {
			p.memory[destIndex] = 1
		} else {
			p.memory[destIndex] = 0
		}
		p.ip += 4
	case AdjustBP:
		amount := p.GetInputOperand(1)
		if p.debug {
			fmt.Printf("bp += %d\n", amount)
		}
		p.bp += amount
		p.ip += 2
	case Halt:
		if p.debug {
			fmt.Printf("HALT\n")
		}
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

func (p *Program) Run() []int {
	for !p.halted {
		p.Step()
	}

	return p.output
}

func (p *Program) RunUntilInput() (halted bool) {
	for p.GetOpcode() != Input && !p.halted {
		p.Step()
	}

	return p.halted
}

/* return output, halted */
func (p *Program) RunUntilOutput() (output int, halted bool) {
	for p.GetOpcode() != Output && !p.halted {
		p.Step()
	}

	// execute the Output opcode (or halt again if we halted; harmless)
	p.Step()

	// return the last output
	return p.output[len(p.output)-1], p.halted
}

/* return opcode, output if output happened, halted */
func (p *Program) RunUntilInputOrOutput() (opcode Opcode, output int, halted bool) {
	for opcode = p.GetOpcode(); opcode != Output && opcode != Input && !p.halted; opcode = p.GetOpcode() {
		p.Step()
	}

	// execute the Output opcode
	if opcode == Output {
		p.Step()
		output = p.output[len(p.output)-1]
	}

	halted = p.halted
	return
}
