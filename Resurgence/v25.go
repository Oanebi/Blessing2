package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func factory(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("code")
	//w.WriteHeader(data)
	if data == "" {
		http.Error(w, "code parameter is required", http.StatusBadRequest)
		return
	}
	result, err := strconv.Atoi(data)
	if err != nil {
		http.Error(w, "code must be a valid integer", http.StatusBadRequest)
		return
	}
	if result < 100 || result > 599 {
		http.Error(w, "code must be a valid HTTP status code (100–599)", http.StatusBadRequest)
		return
	} else {
		w.WriteHeader(result)
		fmt.Fprintf(w, "Responding with status [code]", result)
	}
}
