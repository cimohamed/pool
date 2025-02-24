package main

import (
	"fmt"
)

func ToUpper(s string) string {
	var l string
	for _, i := range s {
		if i >= 'a' && i <= 'z' {
			l += string(i - 32)
		} else {
			l += string(i)
		}
	}
	return l
}
func main() {
	fmt.Println(ToUpper("Hello! How are you?"))
}
