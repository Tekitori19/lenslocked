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
	tpl, err := views.Parse(filepath.Join("templates", "home.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Join("templates", "contact.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/contact", controllers.StaticHandler(tpl))
	r.Get("/contact/{id}", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Join("templates", "faq.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/faq", controllers.StaticHandler(tpl))
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found :", http.StatusNotFound)
	})

	fmt.Println("server start at port 3000")

	http.ListenAndServe(":3001", r)
}
