package main

import (
	"fmt"
)

func BasicAtoi2(s string) int {
	var r int
	tmp := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] >= 'A' && s[i] <= 'z' {
			return 0
		}
		tmp = append(tmp, int(s[i]-48))
	}
	for j := 0; j < len(tmp); j++ {
		r = r*10 + tmp[j]
	}
	return r
}

func main() {
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello"))
}
// $ go run .
// 12345
// 12345
// 0
// 0
// $
