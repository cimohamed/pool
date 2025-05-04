package main

import(
	"fmt"
) 

func LoafOfBread(str string) string {
	words := ""
	word := ""
	for i:= 0 ; i < len(str); {
		if str[i] >='a' && str[i] <= 'z' || str[i] >='A' && str[i] <= 'Z'{
			if len(word) != 5 {
				word += string(str[i])
			}
		}else if  len(word) == 5 {
			words += word
			word = ""
		}
		i++
	}
	return words
}

func main() {
	fmt.Print(LoafOfBread("deliciousbread"))
	fmt.Print(LoafOfBread("This is a loaf of bread"))
	fmt.Print(LoafOfBread("loaf"))
}