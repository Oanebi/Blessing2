package main
import (
	"net/http"
	"fmt"
	"log"
	"strconv"
)
func good(w http.ResponseWriter,r*http.Request){
	if r.Method != http.MethodGet{
		http.Error(w, "use get", http.StatusMethodNotAllowed)
		return
	}
		line := r.URL.Query().Get("weight")
line1 := r.URL.Query().Get("height")
if line == ""{
	http.Error(w, "missing parameter", http.StatusBadRequest)
	return
}
result , err := strconv.ParseFloat(line, 64)
if err != nil{
	fmt.Fprintf(w, "invalid parameter")
	return
}
if line1 == ""{
		http.Error(w, "missing parameter1", http.StatusBadRequest)
		return

}
result1, err := strconv.ParseFloat(line1,64)
if err != nil{
	fmt.Fprintf(w,"invalid parameter1")
	return
}


		bmi := result / (result1*result1)
	
fmt.Fprintf(w,"BMI: %.2f\n",bmi)
	}


func main(){
http.HandleFunc("/bmi",good)
fmt.Println("server is listening at port:8080")
log.Fatal(http.ListenAndServe(":8080",nil))
}