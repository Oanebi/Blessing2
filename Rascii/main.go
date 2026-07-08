package main

import (
	"flag"
	"fmt"
)

func main() {
	package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Dummy Loadbanner function mock for compilation framework context.
// Replace this with your actual local file reading function logic!
func Loadbanner(file string) (map[rune][]string, error) {
	// Your existing banner data loader goes here
	mockMap := make(map[rune][]string)
	return mockMap, nil
}

func main() {
	// Declare the flexible options/flags
	colorFlag := flag.String("color", "", "Color to wrap the text or substring in")
	alignFlag := flag.String("align", "left", "Text alignment (left, center, right, justify)")
	reverseFlag := flag.Bool("reverse", false, "Reverse the output string order")

	flag.Parse()
	remainingArgs := flag.Args()

	if len(remainingArgs) < 1 {
		fmt.Println("Usage: go run . [flags] [substring] <string>")
		return
	}

	bannerfile := "standard.txt"
	colorname := *colorFlag
	alignMode := strings.ToLower(*alignFlag)
	reverseMode := *reverseFlag
	targetsubstring := ""
	input := ""

	// Case A: Just the main string provided
	if len(remainingArgs) == 1 {
		input = remainingArgs[0]
	}

	// Case B: Substring AND main string provided
	if len(remainingArgs) == 2 {
		targetsubstring = remainingArgs[0]
		input = remainingArgs[1]
	}

	banner, err := Loadbanner(bannerfile)
	if err != nil {
		fmt.Println("Error loading banner:", err)
		return
	}

	// Output everything beautifully compiled
	fmt.Print(Render(input, banner, targetsubstring, colorname, alignMode, reverseMode))
}

// 	if len(os.Args) < 2 {
// 		fmt.Println("error")
// 		return
// 	}

// 	input := os.Args[1]
// 	bannerfile := "standard.txt"
// 	colorname := ""
// 	targetsubstring := ""

// 	if len(os.Args) >= 3 {

// 		customfile := os.Args[2]
// 		if !strings.HasSuffix(customfile, ".txt") {
// 			customfile = customfile + ".txt"
// 		}
// 		bannerfile = customfile
// 		if len(os.Args) == 4 {
// 			bannerfile = os.Args[2] + ".txt"
// 			colorname = os.Args[3]

// 		}
// 		if len(os.Args) == 5 {
// 			targetsubstring = os.Args[2]
// 			bannerfile = os.Args[3] + ".txt"
// 			colorname = os.Args[4]
// 		}
// 	}
// 	banner, err := Loadbanner(bannerfile)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Print(Render(input, banner, targetsubstring, colorname))

// }
