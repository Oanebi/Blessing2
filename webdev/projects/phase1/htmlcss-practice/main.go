package main
import (
	"fmt"
	//"log"
	"net/http"
)
// func handler(b http.ResponseWriter, r *http.Request){
// 	fmt.Fprintf(b, "hello from go")
// }
// func main(){
// 	http.HandleFunc("/",handler)
// 	fmt.Println("server running : htp://localhost:8080")
// 	log.Fatal(http.ListenAndServe(":8080",nil)
//}
func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/about", aboutHandler)
    http.HandleFunc("/inspect", inspectHandler)
	http.HandleFunc("/count", TextCounter)
http.HandleFunc("/hello", hellohandler)
	http.HandleFunc("/ping",Handler)


    fmt.Println("server running: http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello from Go!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "This is the about page.")
}

func inspectHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Method: %s\n", r.Method)
    fmt.Fprintf(w, "Path: %s\n", r.URL.Path)
    fmt.Fprintf(w, "Query: %s\n", r.URL.RawQuery)
    fmt.Fprintf(w, "Remote: %s\n", r.RemoteAddr)
	name := r.URL.Query().Get("name")
    age := r.URL.Query().Get("age")
    fmt.Fprintf(w, "Name param: %s\n", name)
    fmt.Fprintf(w, "Age param: %s\n", age)

    for key, vals := range r.Header {
        fmt.Fprintf(w, "Header %s: %s\n", key, vals)
    }
}

