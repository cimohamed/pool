package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	programName := os.Args
	runes := []rune(programName[0])
	for i := range runes {
		if runes[i]=='/' || runes[i]== '.'{
			continue 
		}
		
		z01.PrintRune(runes[i])
	}

	z01.PrintRune('\n')
}
