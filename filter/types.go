package filter

import "github.com/johnsudaar/gitngo/gitprocessor"

// RepositoryStats is a sybtype of the Stats structure returned by the filter
// It hold the number of lines of code written in the filter language present in the current repository
type RepositoryStats struct {
	Repository gitprocessor.GitRepository
	Lines      int
}

// Stats is the structure returned by the Filter Method.
// It hold all the informations extracted from the Github repositories.
type Stats struct {
	Language     string
	Total        int
	Repositories []RepositoryStats
}
