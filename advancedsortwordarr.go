package main

import (
	"fmt"
)

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 0 ; i < len(a) ; i++{
		for j := 0; j < len(a)-1; j++{
			if f(a[j], a[j+1])>0{
				a[j],a[j+1] = a[j+1],a[j]
			}
		}
	} 
}

func Compare(a, b string) int {
	if a > b {
		return 1
	}
	return -1
}
func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	AdvancedSortWordArr(result, Compare)
	fmt.Println(result)
}
