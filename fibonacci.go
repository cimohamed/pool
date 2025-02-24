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
<<<<<<< HEAD
	arg1 := 0
=======
	arg1 := 100
>>>>>>> 39d85a9 (sortparams)
	fmt.Println(fibonacci(arg1))
}
