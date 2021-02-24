package main

import (
	"fmt"
	"html/template"
	"lessons/db"

	"net/http"
	"time"

	"github.com/go-pg/pg"
	"github.com/gorilla/mux"
)

var (
	templates = template.Must(template.ParseGlob("templates/*"))
)


func main() {

	pgdb := db.Connect()
	//SaveProduct(pgdb)
	//DeleteProduct(pgdb)
	UpdateProduct(pgdb)
	//GetIDProduct(pgdb)

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
	templates.ExecuteTemplate(w, "home.html", nil)
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

// SaveProduct is okay
func SaveProduct(dbRef *pg.DB) {
	newPI := &db.ProductItem{
		Name:  "iykecharles316",
		Desc:  "This is my 3rd description",
		Image: "This is my 3rd image",
		Features: struct {
			Name string
			Desc string
			Imp  int
		}{
			Name: "Data engineer",
			Desc: "software architect",
			Imp:  10,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsActive:  true,
	}
	newPI.Save(dbRef)
}

// DeleteProduct is okay
func DeleteProduct(dbRef *pg.DB) {
	newPI := &db.ProductItem{
		Image: "This is my 2nd image",
	}
	newPI.Delete(dbRef)
}

// UpdateProduct is okay
func UpdateProduct(dbRef *pg.DB) {
	newPI := &db.ProductItem{
		Name: "EzemaNew2",
		ID:   1,
	}
	newPI.Update(dbRef)
}

// GetIDProduct is okay
func GetIDProduct(dbRef *pg.DB) {
	newPI := &db.ProductItem{
		ID: 1,
	}
	newPI.GetID(dbRef)
}
