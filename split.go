package main 

import (
	"fmt"
)

func Split(s, sep string)[]string{
	start:= 0
	rus :=[]string{}
	for i:= 0 ; i< len(s)-len(sep);{
		if s[i:i+len(sep)]==sep{
			if i != start{
				rus= append(rus,s[start:i])
			}
			i = i+len(sep)
			start = i
		}else{
			i++
		}
	}
	if start < len(s){
		rus = append(rus,s[start: ])
	}
	return rus
}

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}