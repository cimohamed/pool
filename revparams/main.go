package main

import (
	"fmt"
	"os"
)

func main() {
	for i := len(os.Args) - 1; i > 0; i-- {
		fmt.Println(os.Args[i])
	}
}
