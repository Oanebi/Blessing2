package main
import "fmt"

func MergeBanners(base map[rune][]string, priority map[rune][]string) map[rune][]string {
	result := map[rune][]string{}
	for key, value := range base {
		result[key] = value
	}
	for key, value := range priority {
		result[key] = value
	}
	return result
}
func main() {

	base := map[rune][]string{
		'A': {"A line"},
		'B': {"B line"},
		'C': {"C old"},
	}

	priority := map[rune][]string{
		'B': {"B text"},
		'C': {"C new"},
	}

	result := MergeBanners(base, priority)

	fmt.Println(result)
}
