package main

import (
	"fmt"
)

func WordCount(words []string) map[string]int {
	result := map[string]int{}
	for _, v := range words {
		result[v]++

	}
	return result
}
func main() {
	words := []string{"go", "java", "go", "rust", "java", "go"}
	count := WordCount(words)
	fmt.Println(count)

}
