package main

import (
	"fmt"
)

func StrRev(s string) string {
	var r string
	tmp := []string{}
	for i := len(s) - 1; i >= 0; i-- {
		tmp = append(tmp, string(s[i]))
	}
	for j := 0; j <= len(tmp)-1; j++ {
		r += tmp[j]
	}
	return r

}
func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}
