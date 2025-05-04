package main

import (
	"github.com/01-edu/z01"
)

func hasDuplicateChars(r string) bool {
	for i := 0; i < len(r); i++ {
		for j := i + 1; j < len(r); j++ {
			if r[i] == r[j] {
				return true
			}
		}
	}
	return false
}

func PrintNbrBase(neb int, base string) {
	baseLen := len(base)
	if baseLen < 2 || hasDuplicateChars(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	if neb == 0 {
		z01.PrintRune('0')
		return 
	}
	if neb < 0 {
		z01.PrintRune('-')
		neb = -neb
	}
	ar := []rune{}
	for neb > 0 {
		ruse := neb % baseLen
		ar = append(ar, rune(base[ruse]))
	 	neb/= baseLen
	}
	for i := len(ar) - 1; i >= 0; i-- {
		z01.PrintRune(ar[i])
	}
	return 
}

func main() {
	PrintNbrBase(13, "0123456789")
	z01.PrintRune('\n')
	PrintNbrBase(13, "01")
	z01.PrintRune('\n')
	PrintNbrBase(125, "0123456789ABCDEF")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "choumi")
	z01.PrintRune('\n')
	PrintNbrBase(125, "aa")
	z01.PrintRune('\n')
}

// And its output :
// $ go run .
// 125
// -1111101
// 7D
// -uoi
// NV
// $
