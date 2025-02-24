package main

import (
	"fmt"
)

func AlphaCount(s string) int {
	var cont int
	for i := 0; i < len(s)-1; i++ {
		if (s[i]>='A'&& s[i]<='Z' || s[i]>='a' && s[i]<='z'){
			cont++
		}
		
		
	}
	return cont
}

func main() {
	s := "Hello 78 World!    4455 /"
	nb := AlphaCount(s)
	fmt.Println(nb)
}
