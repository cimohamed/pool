package main

import (
	"fmt"
)
// Index function to find the first occurrence of a substring
func Index(s string, toFind string) int {
		for i:=0 ;i<len(s)-len(toFind);i++{
			if s[i : i+len(toFind)]== toFind{
				return i
			}
		}
		return -1
}

func main() {
	fmt.Println(Index("Hello!", "l"))   
	fmt.Println(Index("Salut!", "u")) 
 	fmt.Println(Index("Ola!", "hOl"))  
}

