package main

import (
	"fmt"
	"io"
	"net/http"
)

func onlypost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "enter a post method", http.StatusMethodNotAllowed)
		return
	}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "error reading the content")
		defer r.Body.Close()
	}
	if len(data) == 0 {
		http.Error(w, "body cannot be empty", http.StatusBadRequest)
		return

	}
	fmt.Fprintf(w, "%s", data)
}
