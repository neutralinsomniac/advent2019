package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("*** PART 2 ***")
	dat, err := ioutil.ReadFile(os.Args[1])
	check(err)

	constructedImage := make([]byte, 150)

	for i := range constructedImage {
		constructedImage[i] = '2'
	}
	for i, c := range dat[:len(dat)-1] {
		if constructedImage[i%150] == '2' && c != '2' {
			constructedImage[i%150] = c
		}
	}

	for i, c := range constructedImage {
		switch c {
		case '0':
			fmt.Printf(" ")
		case '1':
			fmt.Printf("X")
		}
		if (i+1)%25 == 0 {
			fmt.Println("")
		}
	}
}
