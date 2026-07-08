package main
import(
	"net/http"
	"html/template"
)
func switchhandler(w http.ResponseWriter, r *http.Request){
	file1 := r.URL.Query().Get("text")
	file2 := r.URL.Query().Get("banner")
	file3, err := LoadBanner("banner/"+file2)
	if err != nil{
		http.Error(w, "server unable to loadbanner",http.StatusInternalServerError)
		return
	}
	result2 := Render(file1,file3)

	
	tmpl, err := template.ParseFiles("template/index.html")
if err != nil {
	http.Error(w, "unable to parsefile", http.StatusInternalServerError)
	return
}
err = tmpl.Execute(w,pagedata{Result: result2, Text:file1})
if err != nil{
		http.Error(w, "unable to executefile", http.StatusInternalServerError)
return
}

}

