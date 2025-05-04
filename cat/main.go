package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	haserror := false
	if len(os.Args) > 1 {
		for _, filename := range os.Args[1:] {
			file, err := os.Open(filename)
			if err != nil {
				printErure(err)
				haserror = true
				continue
			}
			_, err = io.Copy(os.Stdout, file)
			file.Close()
			if err != nil {
				printErure(err)
				haserror = true
			}

		}
	} else {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			printErure(err)
			haserror = true
		}
	}
	if haserror == true {
		os.Exit(1)
	}
}

func printErure(err error) {
	errr := "ERROR: "
	for _, v := range errr {
		z01.PrintRune(v)
	}
	for _, r := range err.Error() {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
