package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/method-inspector", methodinspector)
	http.HandleFunc("/echo", onlypost)
	http.HandleFunc("/headers", detective)
	http.HandleFunc("/form", decoder)
	http.HandleFunc("/status", factory)

	fmt.Println("listening and serving on port:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
