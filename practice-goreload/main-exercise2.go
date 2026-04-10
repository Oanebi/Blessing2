package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("usage: go run . input.txt output.txt")
		os.Exit(1)
	}
	inputpath := os.Args[1]
	outputpath := os.Args[2]

	inputfile, err := os.Open(inputpath)
	if err != nil {
		fmt.Println("Error opening file", err)

	}
	defer inputfile.Close()
	outputfile, err := os.Create(outputpath)
	if err != nil {
		fmt.Println("Error creating file", err)

	}
	defer outputfile.Close()
	scanner := bufio.NewScanner(inputfile)
	writer := bufio.NewWriter(outputfile)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Fprintf(writer, "%s\n", line)
	}
	writer.Flush()

}
