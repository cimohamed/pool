package main  

import(
	"fmt"
)

func RockAndRoll(n int) string {
	if n < 0 {
		return "error: number is negative\n"
	}
	if n % 2 == 0 && n % 3 == 0 {
		return "rock and roll\n"
	}
	if  n % 2 == 0 || n % 3 == 0 {
		if  n % 2 == 0{
			return "rock\n"
		}else {
			return "roll\n"
		}
	}else{
		return "error: non divisible\n"
	}
	return ""
}

func main() {
	fmt.Println(RockAndRoll(4))
	fmt.Println(RockAndRoll(9))
	fmt.Println(RockAndRoll(6))
	fmt.Println(RockAndRoll(11))
} 