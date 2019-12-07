package main

import (
	"fmt"
	"strconv"
)

func main() {
	count := 0
	for i := 353096; i <= 843212; i++ {
		consec := false
		ascending := true
		str := strconv.Itoa(i)
		for i := range str[:len(str)-1] {
			if str[i] == str[i+1] {
				consec = true
			}
			if str[i+1] < str[i] {
				ascending = false
			}
		}
		if consec && ascending {
			count++
		}
	}
	fmt.Println("*** PART 1 ***")
	fmt.Println("count:", count)
}
