package main

import (
	"fmt"
)

func BasicAtoi(s string) int {
	var r int
	tmp := []int{}
	for i := 0; i < len(s); i++ {
		tmp = append(tmp, int(s[i]-48))
	}
	for j := 0; j < len(tmp); j++ {
		r = r*10 + tmp[j]
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
