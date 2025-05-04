package  main 

import (
	"fmt"
)

func ShoppingSummaryCounter(str string) map[string]int {
	menu := make(map[string]int)
	
	word:=""
	for i := 0 ; i < len(str) ; i++{
		if str[i]==' '{
			if word != ""{
				menu[word]++
				word = ""
			}
		}else{
			word+= string(str[i])
		}
	}
	menu[word]++
	return menu
}

func main() {
	summary := "Burger Water Carrot Coffee Water Water Chips Carrot Carrot Burger Carrot Water"
	for index, element := range ShoppingSummaryCounter(summary) {
		fmt.Println(index, "=>", element)
	}
}