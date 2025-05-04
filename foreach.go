package main

import (
	"github.com/01-edu/z01"
)

func ForEach(f func(int), a []int) {
	for _, v := range a {
		f(v)
	}
	z01.PrintRune('\n')
}

func PrintNbr(n int) {
	z01.PrintRune(rune(n+'0'))
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	ForEach(PrintNbr, a)
}
