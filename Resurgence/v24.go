package main
import(
	"net/http"
	"fmt"
)
func decoder(w http.ResponseWriter, r*http.Request){
if r.Method == http.MethodPost{

r.ParseForm()
data := r.FormValue("username")
result := r.FormValue("language")


if data == ""{
	http.Error(w, "username is required",http.StatusBadRequest)
	return
}
if result == ""{
	http.Error(w,"language is required", http.StatusBadRequest)
	return
}else{
	fmt.Fprintf(w,"Hello %s, you are coding in %s!",data,result)
}

}

http.Error(w, "non-post requests", http.StatusMethodNotAllowed)

} 