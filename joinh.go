package main

import "fmt"

func Join(strs []string, sep string) string {
	word := ""
 	for _,str := range strs{
		word += str+sep
	}
	return word
}
func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(Join(toConcat, ":"))
}