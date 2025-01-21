package main

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
	} else if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	arry := []int{}
	for n > 0 {
		a := n % 10
		arry = append(arry, a)
		n = n / 10
	}
	for i := len(arry) - 1; i >= 0; i-- {
		z01.PrintRune(rune(arry[i] + 48))
	}
}

func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	z01.PrintRune('\n')
}
