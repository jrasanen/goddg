package ddg

import "fmt"

const queryEndpoint = "https://api.duckduckgo.com/?q=%s&format=json&pretty=1&no_redirect=1"

type Icon struct {
	URL    string `json:"URL"`
	Width  string `json:"Width"`
	Height string `json:"Height"`
}

type Result struct {
	Result   string `json:"Result"`
	FirstURL string `json:"FirstURL"`
	Text     string `json:"Text"`
	Icon     string `json:"Icon"`
}

type SearchMeta struct {
	Status          string `json:"status"`
	ProductionState string `json:"production_state"`
	PerlModule      string `json:"perl_module"`
}

type SearchResults struct {
	Heading       string     `json:"heading"`
	AbstractURL   string     `json:"AbstractURL"`
	Meta          SearchMeta `json:"meta"`
	RelatedTopics []Result   `json:"RelatedTopics"`
	Results       []Result   `json:"Results"`
	Redirect      string     `json:"Redirect"`
}

// Query takes a string as parameter and returns SearchResults json
func Query(qs string) *SearchResults {
	query := fmt.Sprintf(queryEndpoint, qs)
	results := new(SearchResults)
	GetJson(query, results)
	return results
}
