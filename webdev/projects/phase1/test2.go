package main
import (
	"net/http"
	"fmt"
	
)
func hellohandler(w http.ResponseWriter, r*http.Request){
	if r.Method != http.MethodGet{
		http.Error(w, "Method not allowed",http.StatusMethodNotAllowed)
	return
	}
	name := r.URL.Query().Get("name")
	if name == ""{
		name = "Guest"
	}
	fmt.Fprintf(w, "Hello, %s!", name)

}
