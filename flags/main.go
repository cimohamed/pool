package main

import (
	"os"

	"github.com/01-edu/z01"
)

func processArgs(args []string) (string, bool, string) {
	var insert, input string
	order := false
	for _, arg := range args {
		if len(arg) >= 9 && hasPrefix(arg, "--insert=") {
			insert = arg[9:]
		} else if len(arg) >= 3 && hasPrefix(arg, "-i=") {
			insert = arg[3:]
		} else if arg == "--order" || arg == "-o" {
			order = true
		} else if arg == "-help" || arg == "-h" {
			printHilp()
			return "", false, ""
		} else {
			input = arg
		}
	}
	return insert, order, input
}
func hasPrefix(arg, prefix string) bool {
	if len(arg) < len(prefix) {
		return false
	}
	for i := 0; i < len(prefix); i++ {
		if arg[i] != prefix[i] {
			return false
		}
	}
	return true
}
func bubbleSort(input string) string {
	runes := []rune(input)
	for i := 0; i < len(runes); i++ {
		for j := 0; j < len(runes)-i-1; j++ {
			if runes[j] > runes[j+1] {
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
		}
	}
	return string(runes)
}

func printHilp() {
	hilps := "--insert\n-i\n	This flag inserts the string into the string passed as argument.\n--order\n-o\n	This flag will behave like a boolean, if it is called it will order the argument.\n"
	for i:= range hilps {
		z01.PrintRune(rune(hilps[i]))
	}
}
func printString(input string){

	for _,v:= range input{
		z01.PrintRune(v)
	} 
	z01.PrintRune('\n')
}

func main() {
	args := os.Args[1:]
	insert, order, input := processArgs(args)

	if input == "" {
		return
	}
	input = input + insert
	if order {
		input = bubbleSort(input)
	}
	printString(input)
}

// $ go run . --insert=4321 --order asdad
// 1234aadds
// $ go run . --insert=4321 asdad
// asdad4321
// $ go run . asdad
// asdad
// $ go run . --order 43a21
// 1234aq
