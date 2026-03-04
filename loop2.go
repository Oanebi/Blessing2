package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter your name:")

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("Enter your age:")
	var age int
	fmt.Scan(&age)
	if age >= 18 {
		fmt.Println("welcome", name, "you are an adult")
	} else {
		fmt.Println("sorry", name, "you are a minor")
	}

}
