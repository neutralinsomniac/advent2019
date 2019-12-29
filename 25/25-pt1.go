package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	combinations "github.com/mxschmitt/golang-combinations"
	"github.com/neutralinsomniac/advent2019/intcode"
)

type Room struct {
	name  string
	exits map[string]*Room
	items []string
}

func ParseRoom(text string) Room {
	room := Room{}
	room.exits = make(map[string]*Room)

	in_doors := false
	in_items := false

	lines := strings.Split(text, "\n")

	for _, s := range lines {
		if len(s) == 0 {
			continue
		}
		if s[0:2] == "==" {
			room.name = s
		}
		if in_doors && s[0] == '-' {
			exit := s[2:]
			room.exits[exit] = nil
			continue
		}
		if in_items && s[0] == '-' {
			item := s[2:]
			room.items = append(room.items, item)
			continue
		}
		if s == "Doors here lead:" {
			in_doors = true
			in_items = false
			continue
		}
		if s == "Items here:" {
			in_items = true
			in_doors = false
			continue
		}
	}
	return room
}

func convertASCIItoIntcode(text string) []int {
	ints := make([]int, len(text))
	for i, c := range text {
		ints[i] = int(c)
	}
	return ints
}

func Explore(program intcode.Program, fromRoom *Room, fromDirection string, seen map[string]bool) *Room {
	newProg := intcode.Program{}
	newProg.InitStateFromProgram(&program)

	newProg.ClearOutput()
	halted := newProg.RunUntilInput()

	if halted {
		fmt.Println("HALTED")
		return nil
	}

	output := newProg.GetASCIIOutput()

	fmt.Println("output", output)
	room := ParseRoom(output)

	switch fromDirection {
	case "north":
		room.exits["south"] = fromRoom
	case "south":
		room.exits["north"] = fromRoom
	case "east":
		room.exits["west"] = fromRoom
	case "west":
		room.exits["east"] = fromRoom
	}

	if seen[room.name] {
		return nil
	}

	seen[room.name] = true

	for direction := range room.exits {
		if room.exits[direction] != nil {
			continue
		}
		ints := convertASCIItoIntcode(direction + "\n")
		newProg2 := intcode.Program{}
		newProg2.InitStateFromProgram(&newProg)
		newProg2.SetReaderFromInts(ints...)
		for range ints {
			newProg2.RunUntilInput()
			newProg2.StepBy(1)
		}
		newRoom := Explore(newProg2, &room, direction, seen)
		if newRoom != nil {
			room.exits[direction] = newRoom
		}
	}
	return &room
}

func printRooms(room *Room) {
	fmt.Println(room)
	for _, r := range room.exits {
		printRooms(r)
	}
}

func main() {
	fmt.Println("*** PART 1 ***")
	program := intcode.Program{}
	program.InitStateFromFile(os.Args[1])

	directions := `north
west
take planetoid
west
take spool of cat6
east
east
south
west
north
take dark matter
south
south
east
west
north
east
east
east
west
north
east
west
west
take coin
north
take jam
south
west
north
west
east
south
south
take wreath
west
take fuel cell
east
north
north
west
drop jam
drop fuel cell
drop planetoid
drop spool of cat6
drop coin
drop dark matter
drop wreath
east
south
east
east
take sand
west
west
north
west
drop sand
`

	manual := false

	inventory := []string{"jam", "fuel cell", "planetoid", "spool of cat6", "coin", "dark matter", "wreath", "sand"}

	program.ClearOutput()
	ints := convertASCIItoIntcode(directions)
	program.SetReaderFromInts(ints...)

	reader := bufio.NewReader(os.Stdin)
	for range ints {
		program.RunUntilInput()
		program.StepBy(1)
	}

	program.RunUntilInput()

	if manual {
		for {
			text, _ := reader.ReadString('\n')
			if strings.Contains(text, "break") {
				break
			}
			ints := convertASCIItoIntcode(text)
			program.SetReaderFromInts(ints...)
			for range ints {
				program.RunUntilInput()
				program.StepBy(1)
			}
			program.RunUntilInput()
			fmt.Println(program.GetASCIIOutput())
		}
	}

	program.ClearOutput()
	for _, items := range combinations.All(inventory) {
		newProg := intcode.Program{}
		newProg.InitStateFromProgram(&program)
		var sb strings.Builder
		for _, item := range items {
			sb.WriteString("take " + item + "\n")
		}
		sb.WriteString("south\n")

		ints := convertASCIItoIntcode(sb.String())
		newProg.SetReaderFromInts(ints...)
		for range ints {
			newProg.RunUntilInput()
			newProg.StepBy(1)
		}
		newProg.RunUntilInput()
		if !strings.Contains(newProg.GetASCIIOutput(), "Alert") {
			fmt.Println(newProg.GetASCIIOutput())
			break
		}
	}
}
