package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i:= 0; i< len(os.Args)-1; i++{
		for j := i ; j<len(os.Args)-i-1; j++{
			if os.Args[j] > os.Args[j+1] {
				os.Args[j],os.Args[j+1] =os.Args[j+1],os.Args[j]
			}
		}
	}
	for _,r := range os.Args[1: ]{
		for _,f := range r{
			z01.PrintRune(f)
			z01.PrintRune('\n')
		}
	} 
}
