package main

import (
	"fmt"
	"strings"
)

func CountVowels(s string) int {
	count := 0
	for _, v := range s {
		if strings.ContainsRune("aeiou", v) {
			count++
		}
	}
	return count

}
func main() {
	fmt.Println(CountVowels("education"))
}
