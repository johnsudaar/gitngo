package webserver

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// This function is reponsible for the Index response
// GET : /
func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello World !\n")
}

// This function is responsible for the Search page response
// GET : /search?query=<QUERY>
func searchHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	query := r.URL.Query().Get("query")
	fmt.Fprintf(w, "Query : %s", query)
}
