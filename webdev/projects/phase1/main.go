package main
import (
	"net/http"
	"fmt"
	"log"
)
func main(){
	http.HandleFunc("/ping", handler)
	http.HandleFunc("/hello", hellohandler)
	http.HandleFunc("/count",TextCounter)
	http.HandleFunc("/calculate", Cal)
	http.HandleFunc("/v2", version)
	http.HandleFunc("/legacy", Legacy)	
		http.HandleFunc("/dashboard", dashboardhandler)
	http.HandleFunc("/agent", User)	


fmt.Println("server listening at port:8080")
log.Fatal(http.ListenAndServe(":8080", nil))

}