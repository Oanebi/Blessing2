package main

import "strings"

func StringToArt(input string) string {
	if input == "" {
		return ""
	}
	patterns := map[rune][]string{

		'1': {
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
		},
		'2': {
			" ___ `",
			"    | ",
			" ___| ",
			"|     ",
			"|___  ",
		},
	}
	for _, char := range input {
		if char == '\n' {
			continue
		}
		if char < '0' || char > '9' {
			return ""
		}
	}
	content := strings.Split(input, "\n")
	var result []string
	for _, word := range content {
		if word == "" {
			continue
		}
		canva := make([]string, 5)
		for _, char := range word {
			pattern, ok := patterns[char]
			if !ok {
				pattern = []string{
					"     ",
					"     ",
					"     ",
					"     ",
					"     ",
				}
			}
			for i := 0; i < 5; i++ {
				canva[i] += pattern[i]

			}
		}
		block := strings.Join(canva, "\n")
		result = append(result, block)

	}
	if len(result) == 0 {
		return ""
	}
	return strings.Join(result, "\n") + "\n"

}
