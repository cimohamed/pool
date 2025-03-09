package main

import (
	"io"
	"os"
	"github.com/01-edu/z01"
)
func main(){
	if len(os.Args) < 3 || os.Args[1] != "-c"{
		return
	}
	n := parseNumber(os.Args[2])
	if n <= 0{
		return
	}
	filse := os.Args[3:]
	if len(filse)==0{
		return
	}
	processFiles(filse, n)
}

func parseNumber(s string)int{
	n:=0
	for _,v:=range s{
		if v < '0' || v > '9'{
			return -1
		}
		n = n*10 + int(v-'0') 
	}
	return n
}
func processFiles(filse []string, n int){
	for i,file:=range filse{
		printFileContent(file,n)
		if i < len(filse)-1{
			z01.PrintRune('\n')
		}
	}
}
func printFileContent(filename string, n int){
	file, err:= os.Open(filename)
	if err != nil{
		return
	}	
	defer file.Close()
	stat,err:= file.Stat()
	if err != nil{
		return
	}
	offset := stat.Size()-int64(n) 
	file.Seek(offset,0)
	io.Copy(os.Stdout,file)
}