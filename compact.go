package main  

import(
	"fmt"
)
const N = 6

func Compact(ptr *[]string) int {
	var arr []string
	for _,v := range *ptr{
		if v != ""{
			arr =append(arr, string(v)) 
		}
	}
	*ptr = arr
	return len(*ptr)
}

func main() {
	a := make([]string, N)
	a[0] = "a"
	a[2] = "b"
	a[4] = "c"

	for _, v := range a {
		fmt.Println(v)
	}

	fmt.Println("Size after compacting:", Compact(&a))

	for _, v := range a {
		fmt.Println(v)
	}
}