package main

import (
	"fmt"
)

func JumpOver(str string) string {
	if len(str) < 3 {
		return "\n"
	}
	result := ""
	for i := 2; i < len(str); i += 3 {
		result += string(str[i])
	}
	result += "\n"
	return result
}

func main() {
	fmt.Print(JumpOver("1010101010"))
	fmt.Print(JumpOver(""))
	fmt.Print(JumpOver("t w l v e"))
	fmt.Print(JumpOver("12"))
}