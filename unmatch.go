package main

import "fmt"

func Unmatch(a []int) int{
	unmach := 0
	for _, char := range a {
		unmach ^=char
	}
	if unmach == 0{
		return -1
	}
	return unmach
}

func main() {
	a := []int{1, 2, 1, 3, 2}
	unmatch := Unmatch(a)
	fmt.Println(unmatch)
}
