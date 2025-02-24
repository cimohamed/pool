package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := 1; i <= len(os.Args)-1; i++ {
		for j := i; j <= len(os.Args)-i-1; j++ {
			if os.Args[j] > os.Args[j+1] {
				os.Args[j], os.Args[j+1] = os.Args[j+1], os.Args[j]
			}
		}
	}

	for _, f := range os.Args[1:] {
		for _, t := range f {
			z01.PrintRune(t)
			z01.PrintRune('\n')
		}

	}

}
