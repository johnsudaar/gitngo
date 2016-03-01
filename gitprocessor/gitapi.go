package gitprocessor

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dghubble/sling"
)

// An hepler function used to construct the client needed to make HTTP request to the github API.
func getSling(path string) *sling.Sling {
	gitBase := sling.New().Base("https://api.github.com/").Client(&http.Client{})
	return gitBase.Get(path)
}

// GetGithubRepositories will get 100 repositories corresponding to the query string sorted by last update date.
func GetGithubRepositories(query string) []GitRepository {
	params := &gitQueryParams{
		Sort:    "updated",
		Query:   query,
		PerPage: 100,
	}
	response := new(gitQueryResponse)
	_, err := getSling("search/repositories").QueryStruct(params).ReceiveSuccess(response)
	if err != nil {
		log.Fatal(err.Error())
	}

	return response.Items
}

// GetRepositoryLanguages will get language statistics for a given repository
// repository : must be the full repository name (ex: docker/docker)
func GetRepositoryLanguages(repository string) *GitLanguages {
	response := new(GitLanguages)
	path := fmt.Sprintf("repos/%s/languages", repository)
	_, err := getSling(path).ReceiveSuccess(response)
	if err != nil {
		log.Fatal(err.Error())
	}
	return response
}
