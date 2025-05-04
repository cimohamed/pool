package main

import "github.com/01-edu/z01"

func DescendComb(){
	for i := 99 ; i >= 0 ; i--{
		for j := i-1 ; j >= 0 ; j--{
			z01.PrintRune(rune(i/10)+'0')
			z01.PrintRune(rune(i%10)+'0')
			z01.PrintRune(' ')
			z01.PrintRune(rune(j/10)+'0')
			z01.PrintRune(rune(j%10)+'0')
			if !(i == 01 && j == 00){
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		} 
	}
	z01.PrintRune('\n')
}

func main() {
	DescendComb()
}