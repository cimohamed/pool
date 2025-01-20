package main

import "github.com/01-edu/z01"

func comb3() {
		for a := 0; a <= 98; a++ {
			for b := a + 1; b <= 99; b++ {
				z01.PrintRune(rune(a/10 + '0'))
				z01.PrintRune(rune(a%10 + '0'))
				z01.PrintRune(' ')
				z01.PrintRune(rune(b/10 + '0'))
				z01.PrintRune(rune(b%10 + '0'))
				if !(a == 98 && b == 99) {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
	func main(){
		comb3()
	}

		
	
