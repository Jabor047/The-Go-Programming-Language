// package github provides a GO API for Github issue tracker
// see https://developer.github.com/v3/search/#search-issues.

package github

import "time"

const IssuesURL = "https://api.github.com/search/issues"

type IssueSearchResult struct{
	TotalCount int `json:"total_count"`
	Items []*Issue
}

type Issue struct{
	Number int
	HTMLURL string `json:"html_url"`
	Title string 
	State string
	User *User
	CreatedAt time.Time `json:"created_at"`
	Body string 		//in markdown format
}
type User struct {
	Login string
	HTMLURL string `json:"html_url"`
}