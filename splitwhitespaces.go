package main

import (
	"fmt"
)

func isWhitespace(char rune) bool {
	return char == ' ' || char == '\n' || char == '\t'
}

func SplitWhiteSpaces(s string) []string {
	rus := []string{}
	word := []rune{}
	for _, char := range s {
		if isWhitespace(char) {
			if len(word) > 0 {
				rus = append(rus, string(word))
				word = nil
			} 
		}else{
			word = append(word, char)
	    }
	}
	if len(word) > 0 {
		rus = append(rus, string(word))
	}
	return rus
}

func main() {
	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
}
