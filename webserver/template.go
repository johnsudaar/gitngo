package webserver

import (
	"log"
	"net/http"
	"text/template"
)

// A simple wrapper for the templating method.
// Will apply layout and manager paths.
func render(w http.ResponseWriter, page string, data interface{}) {
	t := template.New("GitNGo")
	t = template.Must(t.ParseFiles("ressources/html/layout.html.tmpl", "ressources/html/"+page))
	err := t.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Fatal(err.Error())
	}
}
