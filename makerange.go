package main

import (
	"fmt"
)

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	rusulte := make([]int, max-min)
	for i := min; i < max; i++ {
		rusulte[i-min] = i
	}
	return rusulte
}

func main() {
	fmt.Println(MakeRange(5, 10))
	fmt.Println(MakeRange(10, 5))
}
