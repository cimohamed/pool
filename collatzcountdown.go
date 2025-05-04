package main 

import(
	"fmt"
)

func CollatzCountdown(start int) int {
	if start <= 0{
		return -1
	}
	i:=0
	for start != 1{
		if start % 2 == 0 {
			start = start / 2
		}else{ 
			start = start * 3 + 1
		}
		i++
	}
	return i
}

func main() {
	steps := CollatzCountdown(657)
	fmt.Println(steps)
}