package handlers

import (
	"ascii-art-web/ascii"
	"html/template"
	"net/http"
	"strings"
)

// This is the container tray used to pass data to the HTML page
type templateData struct {
	Result    string
	InputText string
	Banner    string
	Error     string
}

func loadTemplate(w http.ResponseWriter) *template.Template {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "500 - Internal Server Error: could not load layout", http.StatusInternalServerError)
		return nil
	}
	return tmpl
}

// HomeHandler serves the initial entry page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "400 - Bad Request: method not allowed", http.StatusBadRequest)
		return
	}

	if r.URL.Path != "/" {
		http.Error(w, "404 - Page Not Found", http.StatusNotFound)
		return
	}

	tmpl := loadTemplate(w)
	if tmpl != nil {
		tmpl.Execute(w, templateData{Banner: "standard"})
	}
}

// AsciiArtHandler accepts the form payload and calls your engine
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "400 - Bad Request: method not allowed", http.StatusBadRequest)
		return
	}

	inputText := r.FormValue("text")
	banner := r.FormValue("banner")

	tmpl := loadTemplate(w)
	if tmpl == nil {
		return
	}

	// 1. Text field can't be empty check
	if strings.TrimSpace(inputText) == "" {
		w.WriteHeader(http.StatusBadRequest)
		tmpl.Execute(w, templateData{Error: "400 - Bad Request: Text input cannot be empty!", Banner: banner})
		return
	}

	// 2. Call your pencil logic to render the artwork matrix
	result, err := ascii.GenerateArt(inputText, banner)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			w.WriteHeader(http.StatusNotFound)
			tmpl.Execute(w, templateData{Error: "404 - Not Found: Banner file missing", InputText: inputText, Banner: banner})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		tmpl.Execute(w, templateData{Error: "500 - Server Error: Machine broke", InputText: inputText, Banner: banner})
		return
	}

	// 3. Render success back onto the browser screen
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, templateData{Result: result, InputText: inputText, Banner: banner})
}
