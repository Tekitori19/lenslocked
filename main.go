package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// tplPath := filepath.Join("templates","home.gohtml")
	tpl, err := template.ParseFiles("templates/home.gohtml")
	if err != nil {
		log.Printf("parsing templates %v", err)
		http.Error(w,"There was an error parsing the template",http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("execute templates %v", err)
		http.Error(w,"There was an error execute the template",http.StatusInternalServerError)
		return
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>this is contact page</h1>")
	fmt.Fprint(w, id)
}

// type Router struct{}

// func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w,r)
// 	case "/contact":
// 		contactHandler(w,r)
// 	default:
// 		http.Error(w,"page not found",http.StatusNotFound)
// 	}
// }

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w,r)
// 	case "/contact":
// 		contactHandler(w,r)
// 	default:
// 		http.Error(w,"page not found",http.StatusNotFound)
// 	}
// }

func main() {
	r := chi.NewRouter()

	r.Get("/",homeHandler)
	r.Get("/contact",contactHandler)
	r.Get("/contact/{id}",contactHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request){
		http.Error(w, "Page not found",http.StatusNotFound)
	})

	fmt.Println("server start at port 3000")

	http.ListenAndServe(":3000",r)
}