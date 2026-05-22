package main

import (
	"fmt"
	"strings"
)

func TrimArtRows(rows []string) []string {

	result := make([]string, len(rows))

	for i, row := range rows {
		result[i] = strings.TrimRight(row, " ")
	}

	return result
}
func main() {
	rows := []string{
		// "  ##  ",
		// " #  # ",
		// " #  # ",
		// " #  # ",
		// " #  # ",
		// " #  # ",
		// " #  # ",
		// "      ",
		"hello ", "world "}
	count := TrimArtRows(rows)
	for _, v := range count {
		fmt.Printf("%q", v)
		fmt.Println()
	}
}
