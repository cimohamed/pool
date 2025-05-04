package main

import (
	"fmt"
)

func Map(f func(int) bool, a []int) []bool {
	arr := []bool{}
	for _, v := range a {
		arr = append(arr, bool(f(v)))
	}
	return arr
}
func IsPrime(n int) bool {
	if n == 1{
		return false
	}
	for i := n - 1; i >= 2; i-- {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result := Map(IsPrime, a)
	fmt.Println(result)
}
