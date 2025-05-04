package main

import (
	"fmt"
	"io"
	"os"
	// "io"
)

func main(){

	if len(os.Args)>2{
		fmt.Print("Too many arguments")
		return
	}else if len(os.Args)<2{
		fmt.Print("File name missing")
		return
	}
	file,err:= os.Open(os.Args[1])
	if err!=nil{
		fmt.Print(err)
	}
	arr:=make([]byte,15)
	file.Read(arr)
	if err != nil && err != io.EOF{
		fmt.Print(err)
	}
	fmt.Print(string(arr))
}