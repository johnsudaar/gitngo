package webserver

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Start is the function used to launch the webserver
// The port argument specifies the port number which will be used by the webserver
// This function will return the error returned by the http.ListenAndServe method
func Start(port int) error {
	router := httprouter.New()
	router.GET("/", indexHandler)
	router.GET("/search", searchHandler)
	return http.ListenAndServe(":"+strconv.Itoa(port), router)
}
