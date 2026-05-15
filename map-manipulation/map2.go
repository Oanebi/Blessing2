package main

import "fmt"

func main() {
	fmt.Println("we have the following foods available")
	foods := map[string]int{
		"Rice":  5,
		"Beans": 10,
		"Oil":   3,
	}
	//how to add to a map
	foods["yams"] = 12
	foods["garri"] = 7

	//updating values
	foods["Rice"] = 20

	// how to delete key a map
	delete(foods, "Oil")

	//loop through a map
	total := 0
	for _, value := range foods {
		total += value
	}
	fmt.Printf("the total value we have is: %d\n", total)

}
