package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type LayerCounts map[byte]int

func main() {
	fmt.Println("*** PART 1 ***")
	dat, err := ioutil.ReadFile(os.Args[1])
	check(err)

	allLayerCounts := make([]LayerCounts, 0)

	for i, c := range dat {
		if i%150 == 0 {
			layer := make(LayerCounts)
			allLayerCounts = append(allLayerCounts, layer)
		}
		allLayerCounts[i/150][c]++
	}
	// find layer with fewest '0' digits
	var smallestLayer LayerCounts
	min := math.MaxInt32
	for _, layer := range allLayerCounts[:len(allLayerCounts)-1] {
		if layer['0'] < min {
			min = layer['0']
			smallestLayer = layer
		}
	}
	fmt.Println(smallestLayer['1'] * smallestLayer['2'])
}
