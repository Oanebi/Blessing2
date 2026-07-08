package main

import (
	
	"net/http"
)

func version(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to version 2"))
}
func Legacy(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/v2", http.StatusMovedPermanently)
}