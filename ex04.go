package main

import "fmt"

func Square(number int) int {
	return number * number
}
func main() {
	var number int
	fmt.Println("Enter a Number:")
	fmt.Scan(&number)
	fmt.Println(Square(number))
}
