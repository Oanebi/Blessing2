package main

import (
	"fmt"
	"net/http"
)

func detective(w http.ResponseWriter, r *http.Request) {
	data := r.Header.Get("X-Custom-Token")
	if data == "" {
		http.Error(w, "X-Custom-Token header is missing", http.StatusBadRequest)
	} else {
		fmt.Fprintf(w, "Token received: %s", data)
	}
	data2 := r.Header.Get("Content-Type")
	if data2 == "" {
		http.Error(w, "Content-Type not provided", http.StatusBadRequest)

	} else {
		fmt.Fprintf(w, "Content-Type %s:", data2)

	}
}
