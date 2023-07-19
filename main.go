package main

import (
	"fmt"
	"lenslocked/controllers"
	"lenslocked/views"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	//parse templete
	//only use MUST in main func
	r.Get("/", controllers.StaticHandler(views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))))

	r.Get("/contact", controllers.StaticHandler(views.Must(views.Parse(filepath.Join("templates", "contact.gohtml")))))
	r.Get("/contact/{id}", controllers.StaticHandler(views.Must(views.Parse(filepath.Join("templates", "conyact.gohtml")))))

	r.Get("/faq", controllers.StaticHandler(views.Must(views.Parse(filepath.Join("templates", "faq.gohtml")))))
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found :", http.StatusNotFound)
	})

	fmt.Println("server start at port 3000")

	http.ListenAndServe(":3001", r)
}
