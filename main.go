package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

/*type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}*/

//type HandlerFunc func(w http.ResponseWriter, r *http.Request)

/*
	func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		f(w, r)
	}

	func pathHandler(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			homeHandler(w, r)
		case "/contact":
			contactHandler(w, r)
		default:
			http.Error(w, "Page not found", http.StatusNotFound)
		}
	}
*/

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<h1>FAQ Page</h1>
<ul>
  <li>
    <b>Is there a free version?</b>
    Yes! We offer a free trial for 30 days on any paid plans.
  </li>
  <li>
    <b>What are your support hours?</b>
    We have support staff answering emails 24/7, though response
    times may be a bit slower on weekends.
  </li>
  <li>
    <b>How do I contact support?</b>
    Email us - <a href="mailto:peppelin@gmail.com">peppelin@gmail.com</a>
  </li>
</ul>
`)
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
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page NOT Found", http.StatusNotFound)
	})
	fmt.Println("Starting server on :3000")
	http.ListenAndServe(":3000", r)
}
