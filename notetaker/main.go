package main
import (
	"net/http"
	"fmt"
)
func main(){
	http.HandleFunc("/notes",notetaking)
	fmt.Println("server is listeniing on port 8080 ")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}