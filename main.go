package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/peppelin/lenslocked/controllers"
	"github.com/peppelin/lenslocked/views"
	"net/http"
	"path/filepath"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	tpl := views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.Parse(filepath.Join("templates", "contact.gohtml")))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.Parse(filepath.Join("templates", "faq.gohtml")))
	r.Get("/faq", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page NOT Found", http.StatusNotFound)
	})
	fmt.Println("Starting server on :3000")
	http.ListenAndServe(":3000", r)
}
