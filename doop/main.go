package main

import (
	"os"
	"github.com/01-edu/z01"
)
const maxInt = 9223372036854775807
const minInt = -9223372036854775808

func printNbr(n int){
	if n == 0{
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	a := 0
	ar :=[]rune{}
	for n > 0 {
		a = n%10
		ar = append(ar, rune(a+'0'))
	n/=10
	}
	for j := len(ar)-1 ; j >= 0 ; j--{
		z01.PrintRune(ar[j])
	}
}
func printstr(s string) {
	for _,v := range s{
		z01.PrintRune(v)
	}
}

func Atoi(s string) (int, bool) {
	if s == ""{
		return 0, false
	}
	sign := 1
	start := 0
	if s[0] == '-'{
		sign = -1
		start = 1
	}
	l := 0
	for  i:=start ; i < len(s); i++{
		if s[i] < '0' || s[i] > '9'{
			return 0 , false
		}
		l=l*10 + int(s[i]-'0')
	}
	return sign * l, true
}

func main() {
	if len(os.Args) != 4 {
		return 
	}
	a, ok0 := Atoi(os.Args[1])
	b, ok1 := Atoi(os.Args[3])
	op := os.Args[2]
	if !ok0 || !ok1{
		return
	}
	if op == "-"{
		printNbr(a-b)
	}else if op == "+"{
		printNbr(a+b)
	}else if op == "*"{
		printNbr(a*b)
	}else if op == "/"{
		if b == 0{
			printstr("no division by 0\n")
		}else{
			printNbr(a/b)
		}
	}else if op == "%"{
		if b == 0{
			printstr("no division by 0\n")
		}else {
			printNbr(a%b)
		}
	}else{
		return
	}
	z01.PrintRune('\n')

}
