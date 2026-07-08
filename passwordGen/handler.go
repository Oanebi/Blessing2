package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"strings"
)

func passwordhandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, err := template.ParseFiles("template/index.html")
		if err != nil {
			http.Error(w, "error parsing files", http.StatusInternalServerError)
			return
		}
		err = data.Execute(w, nil)
		if err != nil {
			http.Error(w, "error executing files", http.StatusInternalServerError)
			return

		}

	} else if r.Method == http.MethodPost {
		// email := r.FormValue("email")
		// if email == "" {
		// 	http.Error(w, "empty email", http.StatusBadRequest)
		// 	return

		numbers := r.FormValue("numbers")
		uppercase := r.FormValue("uppercase")
		lowercase := r.FormValue("lowercase")
		special := r.FormValue("special")

		pool := ""
		if numbers == "on" {
			pool += "0123456789"
		}
		if uppercase == "on" {
			pool += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		}
		if lowercase == "on" {
			pool += "abcdefghijklmnopqrstuvwxyz"
		}
		if special == "on" {
			pool += "!£$%^&*?"
		}
		if pool == "" {
			http.Error(w, "choose a checkbox", http.StatusBadRequest)
			return
		}
		var output strings.Builder
		for i := 0; i < 8; i++ {
			index := rand.Intn(len(pool))
			//index = rand.Intn(len(pool))
			output.WriteString(string(pool[index]))
		}
		password := output.String()
		data, err := template.ParseFiles("template/index.html")
		if err != nil {
			http.Error(w, "error parsing", http.StatusInternalServerError)
			return
		}
		err = data.Execute(w, struct{ Result string }{password})
		if err != nil {
			http.Error(w, "eoor executing", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
	}
}

// from := os.Getenv("GMAIL")
// appPassword := os.Getenv("APP_PASSWORD")
// to := []string{email}
// msg := []byte("Subject: Your Password\n\nYour generated password is: " + password)

// err = smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, appPassword, "smtp.gmail.com"), from, to, msg)
// if err != nil {
// 	fmt.Println("smtp error:", err)
// 	http.Error(w, "error", http.StatusInternalServerError)
// 	return
// }
