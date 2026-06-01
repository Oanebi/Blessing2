package main

import (
	"fmt"
	"regexp"
)

// func CountAlpha(s string) int {
// 	count := 0
// 	for _, v := range s {
// 		if v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z' {
// 			count++
// 		}
// 	}
// 	return count

// }

func main() {
	fmt.Println(CountAlpha("Hello world"))
	fmt.Println(CountAlpha("H e l l o"))
	fmt.Println(CountAlpha("H1e2l3l4o"))
}

// func CountAlpha(s string) int {
// 	count := 0
// 	alpha := strings.Fields(s)
// 	for _, z := range alpha {
// 		for _, w := range z {
// 			if (w >= 'a' && w <= 'z') || (w >= 'A' && w <= 'Z') {
// 				//	if strings.ContainsAny(z){
// 				count++
// 			}
// 		}
// 	}
// 	return count
// }

func CountAlpha(s string) int {
	pattern := regexp.MustCompile(`[a-zA-Z]`)
	matches := pattern.FindAllString(s, -1)
	return len(matches)
}
