package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func parseAndExecuteTpl(w http.ResponseWriter, page string, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles(fmt.Sprintf("templates/%s.gohtml", page))

	if err != nil {
		log.Printf("parsing templates %v", err)
		http.Error(w, "There was an error parsing the template", http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, data)
	if err != nil {
		log.Printf("execute templates %v", err)
		http.Error(w, "there was an error execute the template", http.StatusInternalServerError)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// tplPath := filepath.Join("templates","home.gohtml")
	parseAndExecuteTpl(w, "home", nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	parseAndExecuteTpl(w, "contact", struct{ ID string }{ID: id})
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	parseAndExecuteTpl(w, "faq", nil)
}

func main() {
	r := chi.NewRouter()

	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/contact/{id}", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("server start at port 3000")

	http.ListenAndServe(":3000", r)
}
