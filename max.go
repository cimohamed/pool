package main 

import (
	"fmt"
)

func Max(a []int) int {
	for i := 0; i < len(a); i++{
		for j:= i; j < len(a)-1 ; j++{
			if a[j]>a[j+1]{
				a[j],a[j+1]=a[j+1],a[j]
			}
		}
	} 
	return a[len(a)-1]
}

func main() {
	a := []int{23, 123, 1, 11, 55, 93}
	max := Max(a)
	fmt.Println(max)
}
