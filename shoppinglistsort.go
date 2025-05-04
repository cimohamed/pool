package main

import (
	"fmt"

)

func ShoppingListSort(slice []string) []string {
	for i:= 0 ; i <len(slice); i++{
		for j := 0 ; j < len(slice)-1; j++{
			if len(slice[j])>len(slice[j+1]){
				slice[j],slice[j+1]=slice[j+1],slice[j]
			}
		}
	}
	return slice
}

func main() {
	slice := []string{"Pineapple", "Honey", "Mushroom", "Tea", "Pepper", "Milk"}
	fmt.Println(ShoppingListSort(slice))
}