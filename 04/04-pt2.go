package main

import (
	"fmt"
	"strconv"
)

func main() {
	count := 0

	lower := 353096
	higher := 843212

nextNumber:
	for i := lower; i <= higher; i++ {
		str := strconv.Itoa(i)
		str = str + "f"
		for i := range str[:6] {
			if str[i+1] < str[i] {
				//fmt.Println("fail ascending:", str)
				continue nextNumber
			}
		}

		groupFound := false
		fails := make(map[byte]bool)
		for i := range str[:5] {
			if fails[str[i]] {
				continue
			}
			if str[i] == str[i+2] {
				fails[str[i]] = true
				continue
			}
			if str[i] == str[i+1] {
				groupFound = true
			}

		}

		if !groupFound {
			//fmt.Println("fail no groups:", str)
			continue
		}

		count++
		fmt.Println(i)
	}
	fmt.Println("count:", count)
}
