package main

import (
	"fmt"
)

func ToLower(s string) string {
	var l string
	for _, i := range s {
		if i >= 'A' && i <= 'Z' {
			l += string(i + 32)
		} else {
			l += string(i)
		}
	}
	return l
}
func main() {
	fmt.Println(ToLower("Hello! How are you?"))
}
