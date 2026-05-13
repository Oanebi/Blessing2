package main

import "strings"

func Renderfile(input string, banner map[rune][]string) []string {
	graph := []string{}
	var output strings.Builder
	for i := 0; i < 8; i++ {
		for _, char := range input {
			output.WriteString(banner[char][i])
		}
		graph = append(graph, output.String())
		output.Reset()

	}
	return graph
}
