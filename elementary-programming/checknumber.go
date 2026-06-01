package main

import (
	"fmt"
	//"strings"
)

// func CheckNumber(arg string) bool {
// 	//for _, b := range arg {
// 	numbers := "0123456789"
// 	return strings.ContainsAny(arg, numbers)
// 	//	return true

// }

// //return false

func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
}

func CheckNumber(arg string) bool {
	for _, b := range arg {
		if b >= '0' && b <= '9' {
			return true
		}
	}
	return false

}
