package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 234
	ptr.y = 21
}

func printnbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	var l int
	dig := []rune{}
	for n > 0 {
		l = n % 10
		dig = append(dig, rune(l)+'0')
		n/= 10
	}
	for j := len(dig) - 1; j >= 0; j-- {
		z01.PrintRune(dig[j])
	}
}

func printchars(s string) {
	for _,v:= range s{
		z01.PrintRune(v)
	}
}

func main() {
	points := &point{}
	setPoint(points)
	printchars("x = ")
	printnbr(points.x)
	printchars( ", y = ")
	printnbr(points.y)
	z01.PrintRune('\n')
}
