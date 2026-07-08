package main
import (
	"net/http"
	"fmt"
	
	
)
func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "pong")
}
