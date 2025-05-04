package main

import (
	"fmt"
)

func contalpha(s string) int {
	C := 0
	for _, v := range s {
		if v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z' {
			C++
		}
	}
	return C
}

func main() {
	fmt.Print(contalpha(("hollo world!")))
}
