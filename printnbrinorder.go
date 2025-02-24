package main

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	ar := []rune{}
	for n > 0 {
		ar = append(ar, rune(n%10+'0'))
		n /= 10
	}
	// Bubble sort
	for i := 0; i < len(ar); i++ {
		for j := 0; j < len(ar)-i-1; j++ {
			if ar[j] > ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
			}
		}
	}

	for _, digit := range ar {
		z01.PrintRune(digit)
	}
}
func main() {
	PrintNbrInOrder(321)
	PrintNbrInOrder(0)
	PrintNbrInOrder(391)
}
