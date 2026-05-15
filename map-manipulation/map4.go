package main

import "fmt"

func Highest(foods map[string]int) (string, int) {
	highest := 0
	item := ""
	for key, value := range foods {
		if value > highest {
			highest = value
			item = key
		}
	}
	return item, highest
}
func Lowest(foods map[string]int) (string, int) {
	lowest := foods["Oil"]
	item := "Oil"
	for key, value := range foods {
		if value < lowest {
			lowest = value
			item = key
		}
	}
	return item, lowest
}
func main() {
	foods := map[string]int{
		"Rice":  20,
		"Beans": 10,
		"garri": 7,
		"Oil":   12,
	}
	highItem, highValue := Highest(foods)
	lowItem, lowValue := Lowest(foods)

	fmt.Println("Highest:", highItem, highValue)
	fmt.Println("Lowest:", lowItem, lowValue)
}
