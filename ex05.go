package main

import "fmt"

func Checkpass(score int) string {
	if score <= 50 {
		return "you failed!"
	}
	return "you failed!"
}

func main() {
	fmt.Println("Enter your score:")
	var score int
	fmt.Scan(&score)
	fmt.Println(Checkpass(score))
}
