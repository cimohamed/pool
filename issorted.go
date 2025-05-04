package main 

import(
	"fmt"
)

func f(n, b int) int {
	if n < b {
		return 1
	}
	return -1
}

func IsSorted(f func(a, b int) int, a[]int) bool{
	for i :=0 ; i < len(a)-1; i++{
		if f(a[i], a[i+1]) != 1{
			return false
		}
	}
	return true
}


func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 1, 2, 3, 5,4}

	result1 := IsSorted(f, a1)
	result2 := IsSorted(f, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}