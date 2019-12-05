package main

import (
	"fmt"
	"os"
	"bufio"
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
		total_fuel := fuel
		for (fuel > 0) {
			fuel = calcFuel(fuel)
			if (fuel > 0) {
				total_fuel += fuel
			}
		}
		sum += total_fuel
	}
	fmt.Println(sum)
}
