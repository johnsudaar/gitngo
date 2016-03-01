package main

import (
	"log"

	"github.com/johnsudaar/gitngo/gitprocessor"
	"github.com/johnsudaar/gitngo/webserver"
)

func main() {
	log.Println(len(gitprocessor.GetGithubRepositories("docker")))

	log.Fatal(webserver.Start(8080))
}
