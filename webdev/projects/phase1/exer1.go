//  Register a /method-inspector handler using http.HandleFunc.
// ●     Read the request method using r.Method.
// ●     Respond with a plain text message that includes the method name.
// ○     GET request → "You made a GET request."
// ○     POST request → "You made a POST request."
// ○     Any other method → "You made a [METHOD] request."
// ●     Do not hardcode each method with its own if/else branch — use the value of r.Method directly in your response string.

package main

import (
	"fmt"
	"log"
	"net/http"
)

func methodinspector(w http.ResponseWriter, r *http.Request) {
	//	r.Method = http.MethodGet
	fmt.Fprintf(w, "You made a %s request", r.Method)

}

func main() {
	http.HandleFunc("/method-inspector", methodinspector)
	fmt.Println("server running at port:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
