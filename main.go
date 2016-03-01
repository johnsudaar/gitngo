package main

import (
	"flag"
	"log"

	"github.com/johnsudaar/gitngo/webserver"
)

var portNumber int

func init() {
	flag.IntVar(&portNumber, "port", 8080, "Listenning port")
	flag.Parse()
}

func main() {
	log.Fatal(webserver.Start(portNumber))
}
