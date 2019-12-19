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

func generateFFT(index int, length int) []int {
	baseFFT := []int{0, 1, 0, -1}
	newFFT := make([]int, length+1)
	for i := 0; i < length+1; i++ {
		newFFT[i] = baseFFT[(i/(index+1))%len(baseFFT)]
	}

	// shift the FFT
	copy(newFFT, newFFT[1:])
	return newFFT
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println("*** PART 2 ***")
	var num []int

	dat, err := ioutil.ReadFile(os.Args[1])
	check(err)

	for i := 0; i < 10000; i++ {
		for _, c := range dat[:len(dat)-1] {
			digit, err := strconv.Atoi(string(c))
			check(err)
			num = append(num, digit)
		}
	}

	offset, err := strconv.Atoi(string(dat[:7]))
	check(err)
	// since we're in the second half of the number, and we work backwards from the end of the number, we can discard everything before our relevant offset for a slight speed boost
	num = num[offset:]
	input := num
	for phase := 0; phase < 100; phase++ {
		newInput := make([]int, len(num))
		sum := 0
		for i := len(num) - 1; i >= 0; i-- {
			newInput[i] = (sum + input[i]) % 10
			sum = newInput[i]
		}
		//fmt.Println(newInput)
		input = newInput
	}
	fmt.Println(input[:8])
}
