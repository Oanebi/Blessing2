package main

import (
	"fmt"
	"strings"
)

func Reverse(s string) string {
	var result strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		result.WriteString(string(s[i]))

	}
	return result.String()
}

func main() {
	//Test case:

	fmt.Println(Reverse("golang"))
}

//Expected output:

//gnalog
