package main

import (
	"fmt"
)

func StrLen(s string) int {
	return len(s)
}
func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}
