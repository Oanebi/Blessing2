package main

import "fmt"

func main() {
	fmt.Println("Enter your age:")
	var age int
	fmt.Scan(&age)
	if age >= 18 {
		fmt.Println("you are an adult")
	} else {
		fmt.Println("you are a minor")
	}
}
