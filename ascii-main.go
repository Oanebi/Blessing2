package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . <bannerfile>")
		fmt.Println("       go run . text bannerfile")
		return
	}

	// 2 args — just load and print all characters
	if len(os.Args) == 2 {
		bannerFile := os.Args[1] + ".txt"
		banner, err := LoadBanner(bannerFile)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		for i := ' '; i <= '~'; i++ {
			fmt.Println(i)
			for _, row := range banner[i] {
				fmt.Println(row)
			}
		}
		return
	}

	// 3 args — render specific text
	if len(os.Args) == 3 {
		text := os.Args[1]
		bannerFile := os.Args[2] + ".txt"
		banner, err := LoadBanner(bannerFile)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		output := RenderLine(text, banner)
		fmt.Println(strings.Join(output, "\n"))
	}
}
