package main 

import (
	"github.com/01-edu/z01"
)

func DealAPackOfCards (deck []int) {
	for i := 0 ; i < 4; i++{
		printstr("Player ")
		z01.PrintRune(rune(i+'1'))
		printstr(": ")
		for j := 0 ; j < 3 ; j++{
			num:=deck[i*3+j]
			printNbr(num)
			if j!=2 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func printstr(s string){
	for _,v := range s{
		z01.PrintRune(v)
	}
}

func printNbr(num int){
	arr := []rune{}
	for num > 0 {
		arr = append(arr, rune((num%10)+'0'))
		num/=10
	}
	for j:= len(arr)-1; j >= 0  ; j--{
		z01.PrintRune(arr[j])
	}
}

func main() {
	deck := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	DealAPackOfCards(deck)
}