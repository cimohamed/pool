package main

import (
	"fmt"
)

func StrLen(s string) int {
	var l int
	for i := 0 ;i<=len(s) ;i++{
		 l=i
	}
return l
}
func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}
