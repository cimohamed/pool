package main

import (
	"fmt"
)

func Atoi(s string) int {
    y:=1
	var r int
	tmp := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] >= 'A' && s[i] <= 'z' {
			return 0
		} else if s[0] == '+' && s[1] == '+' {
			return 0
		} else if s[0] == '+' {
			s = s[1:]
		} else if s[0] == '-'&& s[1] == '-'  {
			return 0
		} else if s[0] == '-' {
			y=-y
			s = s[1:]
		}
		tmp = append(tmp, int(s[i]-48))
	}
	for j := 0; j < len(tmp); j++ {
		r = r*10 + tmp[j]		
	}
	return r*y
}

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}

// $ go run .
// 12345
// 12345
// 0
// 0
// $
// package main
// package main

// import (
// 	"fmt"
// )

// func Atoi(s string) int {
// 	y := 1
// 	var r int
// 	tmp := []int{}

// 	if s[0] == '-' {
// 		y = -1
// 		s = s[1:]
// 	} else if s[0] == '+' {
// 		s = s[1:]
// 	}

// 	for i := 0; i < len(s); i++ {
// 		if s[i] < '0' || s[i] > '9' { // التحقق إذا كانت الأحرف غير صالحة
// 			return 0
// 		}
// 		tmp = append(tmp, int(s[i]-'0'))
// 	}

// 	for j := 0; j < len(tmp); j++ {
// 		r = r*10 + tmp[j]
// 	}

// 	return r * y
// }

// func main() {
// 	fmt.Println(Atoi("12345"))         // 12345
// 	fmt.Println(Atoi("0000000012345")) // 12345
// 	fmt.Println(Atoi("012 345"))       // 0
// 	fmt.Println(Atoi("Hello World!"))  // 0
// 	fmt.Println(Atoi("+1234"))         // 1234
// 	fmt.Println(Atoi("-1234"))         // -1234
// 	fmt.Println(Atoi("++1234"))        // 0
// 	fmt.Println(Atoi("--1234"))        // 0
// }
