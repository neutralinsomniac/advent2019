package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println("*** PART 1 ***")
	var num []int

	dat, err := ioutil.ReadFile(os.Args[1])
	check(err)

	for _, c := range dat[:len(dat)-1] {
		digit, err := strconv.Atoi(string(c))
		check(err)
		num = append(num, digit)
	}

	input := num
	baseFFT := []int{0, 1, 0, -1}

	for phase := 0; phase < 100; phase++ {
		newInput := make([]int, len(num))
		for i := 0; i < len(num); i++ {
			for j, digit := range input {
				f := baseFFT[((j+1)/(i+1))%len(baseFFT)]
				newInput[i] += digit * f
			}
			newInput[i] = abs(newInput[i] % 10)
		}
		input = newInput
	}
	fmt.Println(input[:8])
}
