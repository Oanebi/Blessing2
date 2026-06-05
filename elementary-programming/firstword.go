package main

import (
	"fmt"
	"strings"
)

func FirstWord(s string) string {
	if s == "" {
		return "" + "\n"
	}
	word := strings.Fields(s)

	if len(word) == 0 {
		return ""

	}
	return word[0] + "\n"
}

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))
}

// func FirstWord(s string) string {
// 	if s == "" {
// 		return ""
// 	}
// 	word := strings.Fields(s)
// 	if len(word) == 0 {
// 		return ""
// 	}

// 	return word[len(word)-1] + "\n"

// }
