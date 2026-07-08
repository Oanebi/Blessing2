package main
import(
	"net/http"
	"os"
	"encoding/json"
	"html/template"
	"time"
)
type data struct{
	Title string
	Body string
	Date string
	Quote string
	Total string
}
var notes []data

func notetaking(w http.ResponseWriter, r*http.Request){
	if r.Method == http.MethodGet{
data1, err := os.ReadFile("notes.json")
if err != nil{
	http.Error(w, "errors reading the file",http.StatusInternalServerError)
	return
}
 err = json.Unmarshal(data1,&notes)
 if err != nil{
		http.Error(w, "errors storing  the data",http.StatusInternalServerError)
		return

 }
tmpl, err := template.ParseFiles("templates/index.html")
if err != nil{
		http.Error(w, "errors Parsing the file",http.StatusInternalServerError)
return
}
err = tmpl.Execute(w, notes)
if err != nil{
			http.Error(w, "errors Executing the file",http.StatusInternalServerError)
return
}


	}else if r.Method == http.MethodPost{
		err := r.ParseForm()
		if err != nil{
			http.Error(w, "error parsing form",http.StatusInternalServerError)
			return
		}
		title := r.FormValue("title")
		body := r.FormValue("body")
		date := time.Now().Format("2006-01-02 15:04:05")
		quote := r.FormValue("quote")
		total := r.FormValue("total")
notes = append(notes,data{Title:title,Body:body,Date:date,Quote:quote,Total:total})

jsonBytes,err := json.Marshal(notes)
if err != nil{
	http.Error(w, "error2", http.StatusInternalServerError)
	return
}
 err = os.WriteFile("notes.json",jsonBytes,0644)
if err != nil{
	http.Error(w, "error3", http.StatusInternalServerError)
	return
}
http.Redirect(w, r, "/notes", http.StatusSeeOther)

	}else{
		http.Error(w, "invalid methods",http.StatusMethodNotAllowed)
	}
}