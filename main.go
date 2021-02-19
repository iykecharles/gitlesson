package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	templates = template.Must(template.ParseGlob("templates/*"))
)

// Data is good to go
type Data struct {
	Age      int
	Username string
}

func main() {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notfound)
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	http.Handle("/", r)
	//r.HandleFunc("/notfound", notfound)
	fmt.Println("Your website is now online.......")
	http.ListenAndServe(":3000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	data := Data{
		Age:      20,
		Username: "charlesEzema",
	}
	templates.ExecuteTemplate(w, "home.html", data)
}

func contact(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "contact.html", nil)
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	templates.ExecuteTemplate(w, "faq.html", nil)
}

func notfound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	templates.ExecuteTemplate(w, "404page.html", nil)
}
