package main

import (
	"fmt"
)

func BasicAtoi(s string) int {
	var r int
	for j := 0; j < len(s); j++ {
		r =r*10 + int(s[j]-48)
	}
	return r
}

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}

// $ go run .
// 12345
// 12345
// 0
// $
