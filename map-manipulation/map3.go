package main

import "fmt"

func main() {
	//find the highest value
	foods := map[string]int{
		"Beans": 10,
		"Rice":  20,
		"garri": 7,
		"yams":  12,
	}
	highest := foods["Beans"]
	item := ""
	for key, value := range foods {
		if value > highest {
			highest = value
			item = key
		}
	}
	fmt.Println("highest item:", item)
	fmt.Println("value:", highest)

	lowest := foods["Beans"]
	items := ""

	for key, value := range foods {
		if value < lowest {
			lowest = value
			items = key
		}
	}

	fmt.Println("lowest item:", items)
	fmt.Println("value:", lowest)
}

//how to find the lowest value
