package webserver

import (
	"net/http"

	"github.com/johnsudaar/gitngo/filter"
	"github.com/johnsudaar/gitngo/gitprocessor"
	"github.com/julienschmidt/httprouter"
)

// This function is reponsible for the Index response
// GET : /
func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	render(w, "index.html.tmpl", nil)
}

// This function is responsible for the Search page response
// GET : /search?query=<QUERY>
func searchHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	query := r.URL.Query().Get("query")
	language := r.URL.Query().Get("language")
	repositories := gitprocessor.GetGithubRepositories(query)
	stats := filter.Filter(repositories, language)
	render(w, "search.html.tmpl", stats)
}
