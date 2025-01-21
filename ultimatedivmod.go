package main

import (
	"fmt"
)

func UltimateDivMod(a *int, b *int) {
	var c int
	var d int
	c = *a
	d = *b
	*a = c / d
	*b = c % d
}
func main() {
	a := 13
	b := 2
	UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
