package main

import(
	"fmt"
)
func AppendRange(min, max int) []int {
	ar := []int{}
	for i := min ; i < max ;i++{
		ar = append(ar , i)
	}
	return ar
}

func main() {
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(10, 5))
}
