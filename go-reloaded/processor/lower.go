package processor

import (
	"strconv"
	"strings"
)

func Lower(s string) string {
	input := strings.Fields(s)
	for i := 0; i < len(input); i++ {
		if input[i] == "(low)" && i > 0 {
			input[i-1] = strings.ToLower(input[i-1])
			input = append(input[:i], input[i+1:]...)
			i--
		}

		if input[i] == ("(low,") && strings.HasSuffix(input[i+1], ")") {
			result := strings.TrimSuffix(input[i+1], ")")
			num, err := strconv.Atoi(result)
			if err == nil {
				start := i - num
				if start < 0 {
					start = 0
				}
				for j := start; j < i; j++ {
					input[j] = strings.ToLower(input[j])

				}
				input = append(input[:i], input[i+2:]...)
				i--
			}
		}
	}
	return strings.Join(input, " ")

}
