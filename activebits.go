package main

import (
	"fmt"
)

func ActiveBits(n int) int {
	cont := 0
	for n != 0 {
		if n%2 != 0 {
			cont++
		}
		n /= 2
	}
	return cont
}

func main() {
	fmt.Println(ActiveBits(7))
}
