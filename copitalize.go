package main

import (
	"fmt"
)

func isharf(r rune) bool {
	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
		return true
	}
	return false
}

func isra9m(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}

func isharfk(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - 32
	}
	return r
}
func isharfs(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + 32
	}
	return r
}

func Capitaliz(s string) string {
	isk := true
	ar := []rune(s)
	for i := 0; i < len(ar); i++ {
		if isharf(ar[i]) {
			if isk {
				ar[i] = isharfk(ar[i])
			} else {
				ar[i] = isharfs(ar[i])
			}
			isk = false
		} else if isra9m(ar[i]) {
			isk = false
		} else {
			isk = true
		}
	}
	return string(ar)
}

func main() {
	fmt.Print(Capitaliz("hello! How are you? How+are+things+4you?"))
}
