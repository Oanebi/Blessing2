package main

import "fmt"

func main() {
	fmt.Println("Enter a Number:")
	var number int
	fmt.Scan(&number)
	for i := 0; i <= 20; i++ {

		fmt.Println(number, "x", i, "=", number*i)
	}
}
