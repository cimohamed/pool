package main

import (
	"fmt"
)

func IsPrime(nb int) bool {
	if nb == 1 {
		return false
	}
	if nb <= 0 {
		fmt.Println("-1")

	}
	for i := nb - 1; i >= 2; i-- {
		if nb%i == 0 {
			return false

		}

	}

	return true
}

func main() {
	fmt.Println(IsPrime(1))
	fmt.Println(IsPrime(2))
	fmt.Println(IsPrime(3))
	fmt.Println(IsPrime(4))
	fmt.Println(IsPrime(5))
	fmt.Println(IsPrime(5))
}
