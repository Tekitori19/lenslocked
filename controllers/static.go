package controllers

import (
	"lenslocked/views"
	"net/http"
)

type Static struct {
	Template views.Template
}

func (st Static) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	st.Template.Execute(w, nil)
}

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
