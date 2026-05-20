package main

import "fmt"

func TotalFoods(foods map[string]int) int {
	total := 0
	for _, key := range foods {
		total += key

	}
	return total
}
func main() {
	foods := map[string]int{
		"Rice":  20,
		"Beans": 10,
		"Oil":   3,
		"Garri": 7,
	}
	count := TotalFoods(foods)
	fmt.Println(count)
}
