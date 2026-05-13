package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("go run . text bannerfile")
		return
	}
	input := os.Args[1]
	bannerfile := os.Args[2] + ".txt"
	banner, err := LoadBanner(bannerfile)
	if err != nil {
		fmt.Println("loadbanner error:", err)
		return
	}
	result := Generate(input, banner)
	fmt.Print(result)
}
