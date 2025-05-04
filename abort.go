package main

import(
	"fmt"
)

func Abort(a, b, c, d, e int) int {
	ar := []int{a, b, c, d, e}
	for i := 0 ; i < len(ar); i++{
		for j := 0 ; j < len(ar)-i-1 ; j++{
			if ar[j] > ar [j+1]{
				ar[j], ar[j+1] = ar[j+1], ar[j]
			}
		}
	}
	return ar[len(ar) / 2]
}

func main() {
	middle := Abort(1, 10, 2, 13, 14)
	fmt.Println(middle)
}