package main

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	for i := 1; i <= x; i++ {
		for j := 1; j <= y; j++ {
			if i == 1 && j == 1 || i == x && j == 1 || i == 1 && j == y || i == x && j == y {
				z01.PrintRune('o')
			} else if i == 1 && j > 1 && j < y || i == x && j > 1 && j < y {
				z01.PrintRune('-')

			} else if j == 1 && i > 1 && i < x || j == y && i > 1 && i < x {
				z01.PrintRune('|')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
func main() {
	QuadA(3,5)
}
// $ go run .
// o---o
// |   |
// o---o
// $

