package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isUpper() bool {
	return len(os.Args) > 1 && os.Args[1] == "--upper"
}

func isValidNbr(s string) (int, bool) {
	n := 0
	for _, f:= range s{
		if f < '0' || f > '9' {
			return 0, false
		} 
			n = n*10 + int(f-'0')
		}
		if n >= 1 && n <= 26{
		return n , true
		}
	return 0 ,false
}

func printletter(n int ,upper bool){
	var let rune
	if upper {
	let = rune('A'+n-1)
	}else{
		let=rune('a'+n-1)
	}	
	z01.PrintRune(let)
}
func processArguments() {
		 upper:=isUpper()
		if upper{
			os.Args=os.Args[1: ]
		}
		for _,s:= range os.Args[1: ]{
			n,valid:=isValidNbr(s)
			if valid{
				printletter(n,upper)
			}else{
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}

func main() {
	processArguments()
}
