package filter

import "github.com/johnsudaar/gitngo/gitprocessor"

// RepositoryStats is a sybtype of the Stats structure returned by the filter
// It hold the number of lines of code written in the filter language present in the current repository and the proportion of this project written in this language.
type RepositoryStats struct {
	Repository    gitprocessor.GitRepository `json:"repository"`
	Bytes         int                        `json:"bytes"`
	TotalLines    int                        `json:"total_lines"`
	LanguageLines int                        `json:"lines"`
	Percentage    float64                    `json:"percentage"`
}

// Stats is the structure returned by the Filter Method.
// It hold all the informations extracted from the Github repositories.
type Stats struct {
	Language     string            `json:"language"`
	TotalBytes   int               `json:"total_bytes"`
	TotalLines   int               `json:"total_lines"`
	Repositories []RepositoryStats `json:"repositories"`
}

// The Len, Swap and Less methods are here to implement the sort.Interface
// So the Stats object is sortable by number of line.
func (s Stats) Len() int {
	return len(s.Repositories)
}
func (s Stats) Swap(i, j int) {
	s.Repositories[i], s.Repositories[j] = s.Repositories[j], s.Repositories[i]
}
func (s Stats) Less(i, j int) bool {
	return s.Repositories[i].Bytes < s.Repositories[j].Bytes
}
