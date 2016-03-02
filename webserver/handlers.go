package webserver

import (
	"net/http"
	"strconv"

	"github.com/johnsudaar/gitngo/filter"
	"github.com/johnsudaar/gitngo/gitprocessor"
)

// This function is reponsible for the Index response
// GET : /
func indexHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "index.html.tmpl", nil)
}

// This function is responsible for the Search page response
// GET : /search?query=<QUERY>
func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	customQuery := r.URL.Query().Get("custom")
	language := r.URL.Query().Get("language")
	maxRoutinesV := r.URL.Query().Get("max_routines")
	maxRoutines := 10
	if len(language) == 0 {
		http.Redirect(w, r, "/", 307)
	} else {
		if len(customQuery) == 0 || len(query) == 0 {
			query = "stars:>=0"
		}

		if i, err := strconv.Atoi(maxRoutinesV); (err == nil) && (len(customQuery) != 0) {
			maxRoutines = i
		}

		repositories := gitprocessor.GetGithubRepositories(query)
		stats := filter.Filter(repositories, language, maxRoutines)
		render(w, "search.html.tmpl", stats)
	}
}
