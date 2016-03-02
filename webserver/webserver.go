package webserver

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

// RessourcePath store the ressources access path.
var RessourcePath string

// Start is the function used to launch the webserver
// The port argument specifies the port number which will be used by the webserver
// This function will return the error returned by the http.ListenAndServe method
func Start(port int) error {
	log.Println("[WEB] Loading ...")
	// Router configuration
	router := httprouter.New()
	router.GET(logged("/", indexHandler))
	router.GET(logged("/search", searchHandler))

	// Assets configuration
	fs := http.FileServer(http.Dir(RessourcePath + "/assets/"))
	router.NotFound = http.StripPrefix("/assets/", fs)
	log.Println("[WEB] Listenning at : :" + strconv.Itoa(port))
	return http.ListenAndServe(":"+strconv.Itoa(port), router)
}

// Make httprouter compatible with standard http and make it compatible with alice.
func logged(p string, h func(http.ResponseWriter, *http.Request)) (string, httprouter.Handle) {
	return p, loggerMiddleware(alice.New(context.ClearHandler).ThenFunc(h))
}

// Middleware for logging request statistics.
func loggerMiddleware(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		context.Set(r, "params", ps)
		start := time.Now()
		h.ServeHTTP(w, r)
		log.Printf("[WEB] %s %s in %v", r.Method, r.URL.Path, time.Since(start))
	}
}
