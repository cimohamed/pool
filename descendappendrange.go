package main

import "fmt"

func DescendAppendRange(max, min int) []int {
	arr := make([]int,0)
	if max<min{
		return arr
	}
	for i := max; i > min ;i--{
		arr = append(arr, i)
	}
	return arr
}

func main() {
	fmt.Println(DescendAppendRange(10, 5))
	fmt.Println(DescendAppendRange(5, 10))
}