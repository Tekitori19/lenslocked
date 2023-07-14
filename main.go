package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w,"<h1>ADA WONG HOME PAGE</h1>")
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
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Get("/",homeHandler)
	r.Get("/contact",contactHandler)
	r.Get("/contact/{id}",contactHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request){
		http.Error(w, "Page not found",http.StatusNotFound)
	})

	fmt.Println("server start at port 3000")

	http.ListenAndServe(":3000",r)
}