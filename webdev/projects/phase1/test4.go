package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func Cal(w http.ResponseWriter, r *http.Request) {
	operator := r.URL.Query().Get("op")
	value1 := r.URL.Query().Get("a")
	value2 := r.URL.Query().Get("b")

	result, err := strconv.Atoi(value1)
	if err != nil {
		http.Error(w, "error converting", http.StatusBadRequest)
		return
	}
	result1, err := strconv.Atoi(value2)
	if err != nil {
		http.Error(w, "error converting1", http.StatusBadRequest)
		return
	}
	if operator == "add" {
		fmt.Fprintf(w, "Result: %d\n", result+result1)
	} else if operator == "subtract" {
		fmt.Fprintf(w, "Result: %d\n", result-result1)
	} else if operator == "multiply" {
		fmt.Fprintf(w, "Result: %d\n", result*result1)
	} else {
		http.Error(w, "unknown operation", http.StatusBadRequest)
	}

}