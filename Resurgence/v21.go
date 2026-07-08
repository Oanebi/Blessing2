package main

import (
	"fmt"
	"net/http"
)

func methodinspector(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "you made a %s request", r.Method)
}
