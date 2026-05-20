// // // package main

// // // import "fmt"

// // // func main() {
// // // 	foods := map[string]int{
// // // 		"Rice":  20,
// // // 		"Beans": 10,
// // // 		"Oil":   3,
// // // 	}

// // // 	fmt.Printf("%-9s %s\n", "FOOD", "QUANTITY")
// // // 	fmt.Println("-------------------")

// // // 	for k, v := range foods {
// // // 		fmt.Printf("\033[33m%-10s %d\033[0m\n", k, v)
// // // 			//fmt.Printf("\033[34m%-10s %d\033[0m\n", foods)

// // // }
// // // }

// // package main

// // import "fmt"

// // func main() {
// // 	foods := map[string]int{
// // 		"Rice":  20,
// // 		"Beans": 10,
// // 		"Oil":   3,
// // 		"Garri": 7,
// // 	}

// // 	fmt.Printf("%-10s %-10s\n", "FOOD", "QUANTITY")
// // 	fmt.Println("------------------------")

// // 	for k, v := range foods {
// // 		fmt.Printf("%-10s %-10d\n", k, v)
// // 	}
// // }

// package main

// import "fmt"

// // colour helpers
// func Green(text string) string {
// 	return "\033[32m" + text + "\033[0m"
// }

// func Yellow(text string) string {
// 	return "\033[33m" + text + "\033[0m"
// }

// func Red(text string) string {
// 	return "\033[31m" + text + "\033[0m"
// }

// func main() {
// 	foods := map[string]int{
// 		"Rice":  20,
// 		"Beans": 10,
// 		"Oil":   3,
// 		"Garri": 7,
// 	}

// 	fmt.Println(Green(fmt.Sprintf("%-10s %-10s", "FOOD", "QUANTITY")))
// 	fmt.Println("------------------------")

// 	for k, v := range foods {
// 		fmt.Println(Red(fmt.Sprintf("%-10s %-10d", k, v)))
// 	}
// }

package main

import "fmt"

func ColorByValue(value int) string {
	if value >= 15 {
		return "\033[32m%d\033[0m" // green
	} else if value >= 5 {
		return "\033[33m%d\033[0m" // yellow
	}
	return "\033[31m%d\033[0m" // red
}

func main() {
	foods := map[string]int{
		"Rice":  20,
		"Beans": 10,
		"Oil":   3,
		"Garri": 7,
	}

	fmt.Printf("%-10s %-10s\n", "FOOD", "QUANTITY")
	fmt.Println("------------------------")

	for k, v := range foods {
		colored := fmt.Sprintf(ColorByValue(v), v)
		fmt.Printf("%-10s %-10s\n", k, colored)
	}
}
