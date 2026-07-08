package main

import (
	"fmt"
	"io"
	"net/http"
)

func route(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "Send me a POST, PUT, or DELETE")
	case http.MethodPost:
		data, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		} else {
			fmt.Fprintf(w, "Created: %s", string(data))

		}
	case http.MethodPut:
		data1, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "method not allowed1", http.StatusMethodNotAllowed)
			return
		} else {
			fmt.Fprintf(w, "Updated: %s", string(data1))
		}
	case http.MethodDelete:
		fmt.Fprintf(w, "Deleted")

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/echo", route)
	fmt.Println("server running on port: 8080")
	http.ListenAndServe(":8080", nil)
}
