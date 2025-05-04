package main

import (
	"fmt"
)

func ConcatParams(args []string) string {
	rusl := ""
	for i,arg:= range args{
		rusl += arg
		if i != len(args)-1{
			rusl += "\n"
		}
	}
	return rusl
}

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams(test))
}
