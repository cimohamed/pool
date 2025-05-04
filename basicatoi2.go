package main

import (
	"fmt"
)

func BasicAtoi2(s string) int {
	var r int
	for j := 0; j < len(s); j++ {
		if s[j] >= '0' && s[j] <= '9'{
			r = r*10 + int(s[j]-'0')
		}else{
			return 0
		}
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
