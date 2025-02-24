package main

import "fmt"

func FindNextPrime(nb int) int {
	if ftr(nb) == true {
		return nb
	}
	for ftr(nb) == false {
		nb++
	}
	return nb
}
func ftr(nb int) bool {
	if nb == 1 {
		return false
	}
	if nb <= 0 {
		fmt.Println("-1")
	}
	for i := nb - 1; i >= 2; i-- {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(FindNextPrime(14))
	
	fmt.Println(FindNextPrime(4))
}
