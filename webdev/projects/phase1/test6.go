package main

import (
	"fmt"
	"net/http"
)

func dashboardhandler(w http.ResponseWriter, r *http.Request) {
	key := r.Header.Get("X-API-Key")

	if key != "secret123" {
		http.Error(w, "wrong key", http.StatusUnauthorized)
		return
	}
	fmt.Fprintf(w, "Welcome to the dashboard: %s",key)
}
