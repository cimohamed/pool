package main

import "fmt"

func StringToIntSlice(str string) []int{
	ar:= []int{}
	for _,v := range str {
		ar = append(ar, int(v))
	}
	return ar
}

func main() {
	fmt.Println(StringToIntSlice("A quick brown fox jumps over the lazy dog"))
	fmt.Println(StringToIntSlice("Converted this string into an int"))
	fmt.Println(StringToIntSlice("hello THERE"))
}