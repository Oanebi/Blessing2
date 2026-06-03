package main

import (
	"fmt"
	"log"
	"net/http"

	"ascii-art-web/handlers"
)

func main() {
	// Tie our server route pathways to our specialized handlers
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/ascii-art", handlers.AsciiArtHandler)

	port := ":8080"
	fmt.Println("🚀 Web Server running successfully at http://localhost" + port)

	// Turn the power on!
	log.Fatal(http.ListenAndServe(port, nil))
}
