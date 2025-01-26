package main

import (
	"fmt"
)

func IterativePower(nb int, power int) int {
	if power <= 0 {
		return 0
	}
	r := nb
	for i := 1; i < power; i++ {
		r *= nb
	}
	return r
}

func main() {
	fmt.Println(IterativePower(4, 3))
}
