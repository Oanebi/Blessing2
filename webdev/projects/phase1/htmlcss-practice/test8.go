package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	base := r.URL.Query().Get("base")
	exp := r.URL.Query().Get("exp")
	if base == "" {
		http.Error(w, "base is empty", http.StatusBadRequest)
		return
	}
	if exp == "" {
		exp = "2"
	}
	result1, err := strconv.Atoi(base)
	if err != nil {
		http.Error(w, "error converting1", http.StatusBadRequest)
		return
	}
	result2, err := strconv.Atoi(exp)
	if err != nil {
		http.Error(w, "error converting2", http.StatusBadRequest)
		return
	}
	if result2 < 0 {
		http.Error(w, "less than zero", http.StatusBadRequest)
		return
	}
	result := 1
	for i := 0; i < result2; i++ {
		result = result * result1
	}
	fmt.Fprintf(w, "the result is %d\n:", result)

}
func main() {
	http.HandleFunc("/power", handler)
	fmt.Println("server running on port: 8080")
	http.ListenAndServe(":8080", nil)
}
