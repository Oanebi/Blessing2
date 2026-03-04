package main

import "fmt"

func main() {
	fmt.Println("Enter first Number x:")
	var x int
	fmt.Scan(&x)
	fmt.Println("Enter second number y:")
	var y int
	fmt.Scan(&y)
	fmt.Println("Subtraction", x-y)
	fmt.Println("Addition", x+y)
	fmt.Println("Multiplication", x*y)
	fmt.Println("Division", x/y)

}
