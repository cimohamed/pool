package main

import("fmt")

func BasicJoin(elems []string) string {
	strl:=""
	for i:= 0 ; i<=len(elems)-1;i++{
		strl+=elems[i]

	}
	return strl
}

func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(BasicJoin(elems))
}

// $ go run .
// Hello! How are you?
// $
