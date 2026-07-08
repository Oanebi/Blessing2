package main
import (
	"strings"
)
func Render(input string, banner map[rune][]string)string{
	if input == ""{
		return ""
	}
	if input == "\\n"{
		return "\n"
	}
		var output strings.Builder

	words := strings.Split(input, "\n")
	for _, word := range words{
		if word == ""{
			output.WriteString("\n")
			continue
		}
		for i := 0; i < 8; i++{
			for _, ch := range word{
				if line, exist := banner[ch]; exist && i < len(line){
	output.WriteString(line[i])
				}
			}
			output.WriteString("\n")
		}
	}
	return output.String()
}