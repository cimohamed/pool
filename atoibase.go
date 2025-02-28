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
		ruse = ruse*baseLen + 
		
	}
	return ruse
}

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}
