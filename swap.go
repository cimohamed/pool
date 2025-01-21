package main

import (
	"fmt"
)

func Swap(a *int, b *int) {
	var tmp int
	*a = tmp
	*a = *b
	*b = tmp
}

func main() {
	a := 0
	b := 1
	Swap(&a, &b)
	fmt.Println("a", a)
	fmt.Println("b", b)
}
