package main

import (
	"fmt"
)

func StrRev(s string) string {
	tmp := ""
	for i := len(s) - 1; i >= 0; i-- {
		tmp += string(s[i])
	}
	return tmp
}

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}
