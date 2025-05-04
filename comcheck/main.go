package main  
 
import (
	"os"
	"fmt"
)

func main(){
	args := os.Args[1: ]
	for _,v := range args{
		if v == "01" || v == "galaxy 01" || v == "galaxy" {
			fmt.Println("Alert!!!")
		}
	}
}