package main

import "fmt"

func MostFrequent(words []string) string {
	counts := map[string]int{}
	for _, v := range words {
		counts[v]++
	}
	maxCount := 0
	maxWord := ""
	for word, count := range counts {
		if count > maxCount {
			maxCount = count
			maxWord = word
		}
	}

	return maxWord
}
func main() {
	words := []string{"go", "java", "go", "rust", "java", "go"}
	count := MostFrequent(words)
	fmt.Println(count)
}
