package main

import (
	"fmt"

)

func fibonacci(index int )int {
	if index<0{
		return -1
	}
	if index==0{
		return 0
	}
	return index-1
}

func main() {
	arg1 := 0
	fmt.Println(fibonacci(arg1))
}
