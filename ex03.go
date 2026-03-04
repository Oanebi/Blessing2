package main

import "fmt"

func greeting(name string) string {
	return "Hello" + " " + name +" " + "Hope you are having a great day!"
}
func main() {
	var name string
	fmt.Println("Enter your name:")
	fmt.Scan(&name)
	fmt.Println(greeting(name))
}
