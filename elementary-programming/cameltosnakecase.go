package main

import (
	"fmt"
	"strings"
)

func CamelToSnakeCase(s string) string {
	if len(s) == 0 {
		return ""
	}
	for i := 0; i < len(s); i++ {
		char := s[i]
		if (char < 'A' || char > 'Z') && (char < 'a' || char > 'z') {
			return s
		}
	}
	lastindex := len(s) - 1
	lastbyte := s[lastindex]
	if lastbyte >= 'A' && lastbyte <= 'Z' {
		return s
	}
	for i := 0; i < len(s)-1; i++ {
		currentbyte := s[i]
		nextbyte := s[i+1]
		if (currentbyte >= 'A' && currentbyte <= 'Z') && (nextbyte >= 'A' && nextbyte <= 'Z') {
			return s
		}
	}
	var output strings.Builder
	first := s[0]

	if first >= 'A' && first <= 'Z' {
		output.WriteByte(first + 32)
	} else {
		output.WriteByte(first)
	}
	for i := 1; i < len(s); i++ {
		char := s[i]

		// If it's an uppercase letter, inject the underscore and lowercase the letter
		if char >= 'A' && char <= 'Z' {
			output.WriteByte('_')       // Inject underscore
			output.WriteByte(char + 32) // Inject lowercased letter
		} else {
			// If it's already lowercase, just write it as-is
			output.WriteByte(char)
		}
	}

	// Convert the buffer back into a read-only string
	return output.String()
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
