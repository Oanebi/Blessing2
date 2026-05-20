package main

import (
	"fmt"
)

// check for both key and value
func Isavailable(word map[string]int) bool {
	value, ok := word["Beans"]
	if ok && value == 20 {
		return true
	}
	return false
}

func main() {
	foods := map[string]int{
		"Rice":  20,
		"Beans": 20,
		"Oil":   3,
	}
	count := Isavailable(foods)
	fmt.Println(count)
}

//check for just key

// package main

// import "fmt"

// func IsAvailable(word map[string]int) bool {

// 	_, ok := word["Beans"]

// 	return ok
// }

// func main() {

// 	foods := map[string]int{
// 		"Rice":  20,
// 		"Beans": 10,
// 		"Oil":   3,
// 	}

// 	count := IsAvailable(foods)

// 	fmt.Println(count)
// }
