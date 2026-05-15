package main

import "fmt"

func main() {
	menu := map[string]float64{
		"eggs":    1.75,
		"bacon":   3.22,
		"sausage": 1.89,
	}
	fmt.Println("Dinners Breakfast Menu")
	//for dish, price := range menu {
	//	fmt.Println(dish, price)
	//var price float64
	var total float64
	for _, price := range menu {
		total += price
	}
	fmt.Println(total)

}

//}
