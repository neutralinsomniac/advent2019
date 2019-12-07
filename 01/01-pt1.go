package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calcFuel(mass int) int {
	calc := (mass / 3) - 2
	return calc
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		fuel := calcFuel(mass)
		sum += fuel
	}
	fmt.Println("*** PART 1 ***")
	fmt.Println(sum)
}
