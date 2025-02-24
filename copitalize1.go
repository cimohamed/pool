package main

import "fmt"

func isCap(c rune) bool {
	t := true
	if c >= 'A' && c <= 'Z' {
		t = true
	} else if c >= 'a' && c <= 'z' {
		t = false
	}
	return t
}

func isNotChar(c rune) bool {
	if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
		return true
	}
	return false
}

func capNext(c rune) bool {
	if !isNotChar(c) && !(c >= '0' && c <= '9') {
		return true
	} 
	return false
	
}

func Capitalize(s string) string {
	newStr  := ""
	isCapv := true
	for i := 0; i <= len(s)-1; i++ {
			if  capNext(rune(s[i])) {
				newStr += string(s[i])
				isCapv = true
			}else if !isNotChar(rune(s[i])) {
				newStr += string(s[i])
			} else if isCap(rune(s[i])) && isCapv{
				newStr += string(s[i])
				isCapv = false
			} else if isCap(rune(s[i])) && !isCapv {
				newStr += string(s[i]+32)
			} else if !isCap(rune(s[i])) && isCapv {
				newStr += string(s[i]-32)
				isCapv = false
			} else {
				newStr += string(s[i])
			}
	}
	return newStr
}

func main() {
	fmt.Println(Capitalize("hello! How aRe you? How are3things@you?"))
}
