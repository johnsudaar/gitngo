package filter

import (
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/johnsudaar/gitngo/gitprocessor"
)

// Filter will filter each repository found and count the number of lines of code written in the language passed as parameter.
func Filter(repositories []gitprocessor.GitRepository, language string, maxRoutines int) Stats {

	// Checking if maxRoutines is in a correct range.
	if maxRoutines < 1 || maxRoutines > 100 {
		maxRoutines = 10
	}

	log.Println("[FILTER] Starting with " + strconv.Itoa(maxRoutines) + " routines...")
	// Two channels are made 1 if the language is found in the repository and 1 if the language is not found
	ok := make(chan RepositoryStats, len(repositories))
	failed := make(chan int, len(repositories))

	stats := Stats{
		Language:     language,
		Total:        0,
		Repositories: make([]RepositoryStats, 100), // We do not know the array length in advance. We've make it bigger and we will resize it later.
	}

	// Launching the first subroutines
	for i := 0; i < len(repositories) && i < maxRoutines; i++ {
		go filterWorker(repositories[i], language, ok, failed)
	}

	// curPos store the currentPosition in the result array
	curPos := 0
	for i := 0; i < len(repositories); i++ {
		// If there is something in one of the two channel
		select {
		case stat := <-ok:
			// If the language was found
			stats.Repositories[curPos] = stat
			stats.Total = stats.Total + stats.Repositories[curPos].Bytes
			curPos++
		case <-failed:
			// Else
		}

		// If they are some subroutines which has'nt been run
		// At this point we know that a subroutine has terminated so its safe to re-run one.
		if i+maxRoutines < len(repositories) {
			// Run one
			go filterWorker(repositories[i+maxRoutines], language, ok, failed)
		}
	}
	log.Println("[FILTER] Resizing and sorting...")
	// Resizing the array to the right size.
	stats.Repositories = stats.Repositories[:curPos]

	sort.Sort(sort.Reverse(stats))

	log.Println("[FILTER] Done !")
	return stats
}

// Function used as a subroutine in the Filter method
func filterWorker(repository gitprocessor.GitRepository, language string, ok chan RepositoryStats, failed chan int) {
	repoLanguages := *gitprocessor.GetRepositoryLanguages(repository.FullName)

	exists := false
	totalBytes := 0
	langBytes := 0
	for lang, b := range repoLanguages {
		if strings.ToLower(lang) == strings.ToLower(language) {
			exists = true
			langBytes = b
		}
		totalBytes = totalBytes + b
	}
	// If this repository is using this language
	if exists {
		// Send the correct information in the ok channel
		repo := RepositoryStats{
			Repository: repository,
			Bytes:      langBytes,
			Percentage: 100.0 * float64(langBytes) / float64(totalBytes),
		}
		ok <- repo
	} else {
		// Else send something in the failed channel.
		failed <- 0
	}
}
