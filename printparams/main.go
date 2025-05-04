package main

import (

	"os"
	"github.com/01-edu/z01"
)

func main(){
	for _,i:=range os.Args[1: ]{
		for _,l:= range i{
			z01.PrintRune(l)	
		}
		z01.PrintRune('\n')
	}
}