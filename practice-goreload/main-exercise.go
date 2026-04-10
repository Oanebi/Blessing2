package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cap(s string) string {
	return strings.Title(s)
}
func countwords(s string) int {
	words := strings.Fields(s)
	return len(words)
}
func countchar(s string) int {
	return len([]rune(s))
}
func reverse(s string) string {
	result := ""
	for i := len(s) - 1; i >= 0; i-- {
		result = result + string(s[i])
	}
	return result
}

func main() {
	text := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter: go run input.txt, output.txt")
	for text.Scan() {
		line := text.Text()
		result := strings.Fields(line)
		if len(result) == 2 {
			file1, err := os.Open(result[0])
			if err != nil {
				fmt.Println("error opening file", err)
			}
			defer file1.Close()
			file2, err := os.Create(result[1])
			if err != nil {
				fmt.Println("error creating file", err)
			}
			defer file2.Close()
			writer := bufio.NewWriter(file2)
			scanner := bufio.NewScanner(file1)
			for scanner.Scan() {
				result := scanner.Text()

				line := cap(result)
				words := countwords(result)
				chars := countchar(result)
				reverse := reverse(result)
				fmt.Fprintf(writer, "_____________________\n")
				fmt.Fprintf(writer, "capitalized:%s\n", line)
				fmt.Fprintf(writer, "countwords:%d\n", words)
				fmt.Fprintf(writer, "countchar:%d\n", chars)
				fmt.Fprintf(writer, "reverse:%s\n", reverse)
				fmt.Fprintf(writer, "---------------------- \n")

			}

			writer.Flush()

		} else {
			fmt.Println("enter input.txt, output.txt")
		}
		break

	}

}

