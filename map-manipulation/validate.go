package main
import "fmt"
func ValidateBanner(banner map[rune][]string) error {
	if banner == nil{
		return fmt.errorf("banner is nil", err)
	}
	if len(banner) != 95{
		return fmt.errorf("incomplete")
	}
	for r := 32; r <= 126; r++{
		lines, ok := banner[r]
		if !ok{
return fmt.errorf("invalid")
		}
		if len(lines) != 8{
		return	fmt.errof("incomplete rows")
		}
	}
}
func main() {

	banner := map[rune][]string{
		'A': {"line1","line2","line3","line4","line5","line6","line7","line8"},
		'B': {"1","2","3","4","5","6","7","8"},
	}

	err := ValidateBanner(banner)

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Println("Banner is valid ✅")
}