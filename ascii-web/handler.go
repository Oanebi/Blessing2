package main
import(
	"net/http"
	"html/template"
	
)

type pagedata struct{
	Result string
	Text string
}
func handler(w http.ResponseWriter, r*http.Request){
if r.Method == http.MethodGet{

tmpl, err := template.ParseFiles("template/index.html")
if err != nil{
		http.Error(w, "Error loading template",http.StatusInternalServerError)
return
}
err = tmpl.Execute(w, nil)
if err != nil {
	http.Error(w, "error", http.StatusInternalServerError)
	return
}
}else if 
r.Method == http.MethodPost{

inputstring := r.FormValue("input")
banners1 := r.FormValue("Banners")
banners2, err := LoadBanner("banner/"+banners1)
if err != nil{
	http.Error(w, "Error loading banner",http.StatusInternalServerError)
	return
}
result := Render(inputstring,banners2)

tmpl, err := template.ParseFiles("template/index.html")
if err != nil{
		http.Error(w, "Error loading template",http.StatusInternalServerError)
return
}
 err = tmpl.Execute(w,pagedata{Result: result,Text:inputstring})
 if err != nil{
	http.Error(w, "error acessing result", http.StatusInternalServerError)
	return
}
} else{
	http.Error(w, "Enter a valid Method", http.StatusMethodNotAllowed)
	return
}
}

