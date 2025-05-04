package main

import(
	"fmt"
)

type food struct{
	preptime int
}


func FoodDeliveryTime(order string) int {
	menu := map[string]food{
		"nuggets" : {preptime : 15},
		"burger" : {preptime : 10},
		"chips" : {preptime : 10},
	}
	if item, ok := menu[order]; ok{
		return item.preptime
	}
	return 404
}


func main() {
	fmt.Println(FoodDeliveryTime("burger"))
	fmt.Println(FoodDeliveryTime("chips"))
	fmt.Println(FoodDeliveryTime("nugets"))
	fmt.Println(FoodDeliveryTime("burger") + FoodDeliveryTime("chips") + FoodDeliveryTime("nuggets"))
}