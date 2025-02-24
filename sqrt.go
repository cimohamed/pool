package main

import (
	"fmt"
)

func sqrt(nb int) int {

	var l int
	if nb == 1 {
		return 1
	}
	if nb == 0 {
		return 0
	}
	for i := nb; i > 1; i-- {
		if nb == i*i {
			l = i
		}
	}
	return l
}

func main() {
	fmt.Println(sqrt(81))
	fmt.Println(sqrt(2))
}
