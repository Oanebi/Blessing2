package main
import (
	"net/http"
	"fmt"
)
func User(w http.ResponseWriter, r *http.Request){
	key := r.Header.Get("User-Agent")
	if key == ""{
		http.Error(w, "header not present", http.StatusBadRequest)
	return
	}
	fmt.Fprintf(w,"You are visiting us using: %s\n",key)
}