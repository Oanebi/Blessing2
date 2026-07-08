package main
import(
	"net/http"
	"fmt"
	"log"
)
func main(){
	http.HandleFunc("/ascii-art",handler)
	http.HandleFunc("/ascii-switch",switchhandler)
	fmt.Println("server is listening at port:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}