package processor

import (
	"strconv"
	"strings"
)

func Upper(s string) string {
	input := strings.Fields(s)
	for i := 0; i < len(input); i++ {
		if input[i] == "(up)" && i > 0 {
			input[i-1] = strings.ToUpper(input[i-1])
			input = append(input[:i], input[i+1:]...)
			i--
		}

		if input[i] == ("(up,") && strings.HasSuffix(input[i+1], ")") {
			result := strings.TrimSuffix(input[i+1], ")")
			num, err := strconv.Atoi(result)
			if err == nil {
				start := i - num
				if start < 0 {
					start = 0
				}
				for j := start; j < i; j++ {
					input[j] = strings.ToUpper(input[j])

				}
				input = append(input[:i], input[i+2:]...)
				i--
			}
		}
	}
	return strings.Join(input, " ")

}
