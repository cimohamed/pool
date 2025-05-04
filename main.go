package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]
	for i := 0; i < len(name); i++ {
		if name[i] == '/' {
			name = name[i+1:]

		}
	}
	for i := 0; i < len(name); i++ {
		z01.PrintRune(rune(name[i]))
	}
	z01.PrintRune('\n')
}
