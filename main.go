package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h2>Yes, our website is online</h2>")
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}
