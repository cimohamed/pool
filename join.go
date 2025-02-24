package main

import("fmt")
func Join(strs []string, sep string) string {
	strl:=""
	for i:= 0 ; i<=len(strs)-1;i++{
		strl+=strs[i]+sep

	}
	return strl
}

func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(Join(toConcat, ":"))
}
