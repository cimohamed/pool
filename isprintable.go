package main

import (
	"fmt"
)

func IsPrintable(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] <= 32 || s[i] == 127 {
			return false
		}
	}

	return true
}
func main() {
	fmt.Println(IsPrintable("Hello"))
	fmt.Println(IsPrintable("Hello\n"))

}
