package main

import (
	"fmt"
	"net/http"
)

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		// TODO: set handler for page not found
	}
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Cobtact Page</h1><p>To get in touch, email me at <a href=\"mailto:peppelin@gmail.com\">peppelin@gmail.com</a>")
}
func main() {
	http.HandleFunc("/", pathHandler)
	fmt.Println("Starting server on :3000")
	http.ListenAndServe(":3000", nil)
}
