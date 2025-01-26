package main

import (
	"fmt"
)

func RecursiveFactorial(nb int) int {
	if nb == 1 {
		return 1	
	}
	if nb > 1 {
		return nb * RecursiveFactorial(nb-1)

	}
	return 0
}

func main() {
	arg := 4
	s := RecursiveFactorial(arg)
	fmt.Println(s)
}

// And its output :

// $ go run .
// 24
// $