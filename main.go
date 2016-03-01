package main

import (
	"log"

	"github.com/johnsudaar/gitngo/webserver"
)

func main() {
	log.Fatal(webserver.Start(8080))
}
