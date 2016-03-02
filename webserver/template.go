package webserver

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

// A simple wrapper for the templating method.
// Will apply layout and manager paths and add helper methods.
func render(w http.ResponseWriter, page string, data interface{}) {
	t := template.New("GitNGo")
	t = t.Funcs(template.FuncMap{
		"marshal": func(v interface{}) template.JS {
			a, _ := json.Marshal(v)
			return template.JS(a)
		},
	})
	t = template.Must(t.ParseFiles(RessourcePath+"/html/layout.html.tmpl", "ressources/html/"+page))
	err := t.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Fatal(err.Error())
	}
}
