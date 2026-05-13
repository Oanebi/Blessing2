package main

import "strings"

func Generate(input string, banner map[rune][]string) string {
	if input == "" {
		return ""
	}
	if input == "\\n" {
		return "\n"
	}
	words := SplitInput(input)
	var output strings.Builder
	for i, word := range words {
		if word == "" {
			if i == len(words)-1 {
				for range 8 {
					output.WriteString("\n")
				}

			} else {
				output.WriteString("\n")
			}
			continue
		}
		row := Renderfile(word, banner)
		for _, rows := range row {
			output.WriteString(rows + "\n")
		}

	}
	return output.String()

}
