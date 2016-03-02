package main

import (
	"flag"
	"log"
	"os"

	"github.com/johnsudaar/gitngo/webserver"
)

var portNumber int

func init() {
	flag.IntVar(&portNumber, "port", 8080, "Listenning port")
	flag.StringVar(&webserver.RessourcePath, "ressources", "ressources", "Ressources path")
	flag.Parse()
}

func main() {
	if _, err := os.Stat(webserver.RessourcePath + "/.ressources_flag"); os.IsNotExist(err) {
		log.Fatal("Cannot load ressources.")
	}
	log.Fatal(webserver.Start(portNumber))
}
