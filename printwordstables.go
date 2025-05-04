package main

import (
	"github.com/01-edu/z01"
)

func isWhiteSpace(v rune) bool {
	return v == ' ' || v == '\n' || v == '\t'
}
func SplitWhiteSpaces(a string) []string {
	rus := []string{}
	word := []rune{}
	for _, v := range a {
		if isWhiteSpace(v) {
			if len(word) > 0 {
				rus = append(rus, string(word))
				word = nil
			}
		} else {
			word = append(word, v)
		}
	}
	if len(word) > 0 {
		rus = append(rus, string(word))
	}
	return rus
}
func PrintWordsTables(a []string) {
	for _, v := range a {
		for _, t := range v {
			z01.PrintRune(t)
		}
		z01.PrintRune('\n')
	}
	
}
func main() {
	a := SplitWhiteSpaces("Hello how are you?")
	PrintWordsTables(a)
}
