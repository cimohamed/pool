package main

import (
	"fmt"
<<<<<<< HEAD
=======
	
>>>>>>> 39d85a9 (sortparams)
)

func StrRev(s string) string {
	var r string
<<<<<<< HEAD
	tmp := []string{}
	for i := len(s) - 1; i >= 0; i-- {
		tmp = append(tmp, string(s[i]))
	}
	for j := 0; j <= len(tmp)-1; j++ {
		r += tmp[j]
	}
	return r

=======
	tmp := []rune(s)
	for j:=len(tmp)-1;j>=0;j--{
		r+=string(tmp[j])
	}
return r
>>>>>>> 39d85a9 (sortparams)
}
func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}
