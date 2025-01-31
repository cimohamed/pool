package main

import (
	"fmt"
)

func RecursivePower(nb int, power int) int {
	// if power == 0 {
	// 	return 
	r:=4
var t int
	
	if power >0{

	t= r* RecursivePower(nb, power-1)
	}
	
return t

}


func main() {
	fmt.Println(RecursivePower(4, 3))
}
