package main

import (
	"strings"
)

func Atricle(s string) string {
	input := strings.Fields(s)
	for i := 0; i < len(input); i++ {
		vowels := "aeiouhAEIOUH"
		if input[i] == "a" && strings.Contains(vowels, input[i+1][:1]) {
			input[i] = "an"
		}
		if input[i] == "A" && strings.Contains(vowels, input[i+1][:1]) {
			input[i] = "An"
		}
		if input[i] == "an" && !strings.Contains(vowels, input[i+1][:1]) {
			input[i] = "a"
		}
		if input[i] == "An" && !strings.Contains(vowels, input[i+1][:1]) {
			input[i] = "A"

		}
		if input[i] == "a" && (input[i+1] == "'" || input[i+1] == `"`) && strings.Contains(vowels, input[i+1][:1]) {
			input[i] = "an"
		}
		if input[i] == "A" && (input[i+1] == "'" || input[i+1] == `"`) && strings.Contains(vowels, input[i+1][:1]) {
			input[i] = "An"
		}
	}
	return strings.Join(input, " ")
}
