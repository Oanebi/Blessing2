package main

import (
	"fmt"
	"strings"
)

func PadArtRows(rows []string, width int) []string {

	result := make([]string, len(rows))

	for i, row := range rows {

		if len(row) >= width {
			result[i] = row
		} else {
			padding := width - len(row)
			result[i] = row + strings.Repeat(" ", padding)
		}
	}

	return result
}

func main() {

	rows := []string{
		"hi",
		"hello",
		"a",
		"Go",
	}

	result := PadArtRows(rows, 5)

	for _, r := range result {
		fmt.Println(r, "len:", len(r))
	}
}
