package main

import (
	"fmt"
	"sort"
)

func StackTwo(top []string, bottom []string) []string {

	result := make([]string, len(top)+len(bottom))

	copy(result, top)
	copy(result[len(top):], bottom)
	sort.Strings(result)
	return result
}

func main() {

	top := []string{"A", "B", "C", "F"}
	bottom := []string{"C", "D"}

	result := StackTwo(top, bottom)

	for _, v := range result {
		fmt.Println(v)
	}
}
