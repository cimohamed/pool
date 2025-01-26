
package main

import (
	"fmt"
)

func IterativeFactorial(nb int) int {
	r := 1
	if nb < 0 {
		return 0
	}
	for i := 1; i <= nb; i++ {
		r *= i
	}
	return r
}

func main() {
	arg := 4
	s := IterativeFactorial(arg)
	fmt.Println(s)
}

