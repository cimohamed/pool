package main

import (
	"fmt"
)

func DigitLen(n, base int) int {
	if n < 0 || base < 0{
		fmt.Print('-')
		n = -n
	}
	if n == 0 {
		return 0
	}
	if base < 2 || base > 32 {
		return -1
	}
	a := 0
	for n > 0 {
		if n/base != 0) {
			a++
			n = n / base
		}
	}
	return a
}

func main() {
	fmt.Println(DigitLen(100, 10))
	fmt.Println(DigitLen(100, 2))
	fmt.Println(DigitLen(100, 16))
	fmt.Println(DigitLen(100, 1))
}