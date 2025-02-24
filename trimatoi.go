package main

import (
	"fmt"
)

func cleanCode(str string) (string, bool) {
	cleanstr := ""
	isn := false
	for _, t := range str {
		if t <= '9' && t >= '0' {

			cleanstr += string(t)
		}
		if cleanstr == "" && t == '-' {
			isn = true
		}
	}
	if cleanstr == "" {
		return "errer", false
	}
	return cleanstr, isn
}
func atoi(cleanstr string) int {
	l := 0
	for i := 0; i < len(cleanstr); i++ {
		l = l*10 + int(cleanstr[i]-'0')
	}
	return l
}
func TrimAtoi(s string) int {
	cleanstr, isn := cleanCode(s)

	if cleanstr == "errer" {
		return 0
	}
	vul := atoi(cleanstr)
	if isn == true {
		vul = atoi(cleanstr) * (-1)
	}
	return vul
}
func main() {
	fmt.Println(TrimAtoi("12345"))
	fmt.Println(TrimAtoi("str123ing45"))
	fmt.Println(TrimAtoi("012 345"))
	fmt.Println(TrimAtoi("-Hello World!"))
	fmt.Println(TrimAtoi("sd+x1fa2W3s4"))
	fmt.Println(TrimAtoi("sd-x1fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1-fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1+fa2W3s4"))
}
