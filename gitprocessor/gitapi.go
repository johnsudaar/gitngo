package gitprocessor

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dghubble/sling"
)

// An hepler function used to construct the client needed to make HTTP request to the github API.
func getSling(path string) *sling.Sling {
	gitBase := sling.New().Base("https://api.github.com/").Client(&http.Client{})
	gitBase = gitBase.Set("User-Agent", "gitngo")
	gitKey := os.Getenv("GITHUB_KEY")
	if len(gitKey) != 0 {
		gitBase = gitBase.Set("Authorization", "token "+gitKey)
	}
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

// GetRepositoryLines will get the total lines number of a repository.
// This will use the weekly commit count for a specific repository github API method.
func GetRepositoryLines(repository string) int {
	var response gitCodeFrequency

	path := fmt.Sprintf("repos/%s/stats/code_frequency", repository)

	responseCode := 0
	for tries := 0; tries < 20 && responseCode != 200; tries++ {
		resp, err := getSling(path).ReceiveSuccess(&response)
		responseCode = resp.StatusCode
		// (200 : OK) (202 : Github is computing results)
		if responseCode != 200 && responseCode != 202 {
			log.Fatal(err)
		} else if responseCode == 202 {
			time.Sleep(50 * time.Millisecond)
		}
	}

	// The response is Array of int Array.

	// There is 3 integer per week :
	// [0] : timestamp
	// [1] : addition
	// [2] : deletion
	sum := 0
	for _, week := range response {
		// Additions
		sum += week[1]
		// Deletions
		sum += week[2]
	}

	return sum
}
