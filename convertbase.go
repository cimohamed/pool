package main 

import (
	"fmt"
)

func isValidBase( base string)bool{
	if len(base)<2{
		return false
	}
	sen:=make(map[rune]bool)
	for i:=0 ; i<len(base);i++{
		c:=rune(base[i])
		if sen[c] || c == '+' || c == '-'{
		return false
		}
		sen[c]=true
	}
	return true 
}
func getIndexInBase(c rune , base string)int{
	for i:= 0 ; i< len(base); i++{
		if rune(base[i])==c{
			return i
		}
	}
	return -1
}

func AtoiBase(s string, base string) int {
	if !isValidBase(base){
		return 0
	}
	ruse:=0
	baseLen:= len(base)
	for i:=0 ;i<len(s);i++{
	    index:= getIndexInBase(rune(s[i]), base)
		if index== -1 {
			return 0
		}
		ruse = ruse*baseLen + index
		
	}
	return ruse
}
func Itobase(n int,base string)string{
	if n == 0{
		return string(base[0])
	}
	baseLen := len(base)
	rus := ""
	for n>0{
		rus = string(base[n%baseLen]) + rus
		n/= baseLen
	}
	return rus
}

func ConvertBase(nbr, baseFrom , baseto string)string{
	dec := AtoiBase(nbr, baseFrom)
	return Itobase(dec,baseto)
}

func main() {
	fmt.Println(ConvertBase("101011", "01", "0123456789"))       // 43
	fmt.Println(ConvertBase("7B", "0123456789ABCDEF", "01"))     // 1111011
	fmt.Println(ConvertBase("123", "0123456789", "0123456789ABCDEF")) // 7B
}