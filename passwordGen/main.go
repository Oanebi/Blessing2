package main
import(
	"net/http"
	"fmt"
)
func main(){
	http.HandleFunc("/password", passwordhandler)
	fmt.Println("server is listening on port:8080")
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		fmt.Print(err)
	}
}