package main

import (
	"strconv"
	"strings"
)

func Hextodec(s string) string {
	input := strings.Fields(s)
	for i := 0; i < len(input); i++ {
		if input[i] == "(hex)" && i > 0 {
			num, err := strconv.ParseInt(input[i-1], 16, 64)
			if err == nil {
				input[i-1] = strconv.FormatInt(num, 10)
				input = append(input[:i], input[i+1:]...)
				i--
			}
		}
	}
	return strings.Join(input, " ")
}
