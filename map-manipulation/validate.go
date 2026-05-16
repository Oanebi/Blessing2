package main

import "fmt"

func ValidateBanner(banner map[rune][]string) error {

	if banner == nil {
		return fmt.Errorf("banner is nil")
	}

	if len(banner) != 95 {
		return fmt.Errorf("incomplete")
	}

	for r := rune(32); r <= 126; r++ {

		lines, ok := banner[r]

		if !ok {
			return fmt.Errorf("invalid")
		}

		if len(lines) != 8 {
			return fmt.Errorf("incomplete rows")
		}
	}

	return nil
}

func main() {

	banner := map[rune][]string{
		'A': {"line1", "line2", "line3", "line4", "line5", "line6", "line7", "line8"},
		'B': {"1", "2", "3", "4", "5", "6", "7", "8"},
	}

	err := ValidateBanner(banner)

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Println("Banner is valid ✅")
}