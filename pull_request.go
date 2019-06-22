package main

import (
	"regexp"
)

type PullRequest struct {
	Id      int    `json:"number"`
	HtmlUrl string `json:"html_url"`
	Title   string `json:"title"`
}

type PullRequestList []PullRequest

// Check whether this Pull Request is eligible to be included
// in the sync process by looking at the following conditions:
//
//   - If PR Title starts with '[TP#1234]', case insensitive.
//
// NOTE: The Github API returns a list of "issues", and all pull requests
// are issues (although not all issues are pull requests). Given the limited
// number of fields we parse here, we still blindly model all the responses
// as pull requests anyway. And if we retrieve an issue that matches the
// above conditions, we process it anyway.
//
func (pr *PullRequest) ShouldSync() bool {

	re := regexp.MustCompile(`^(\[TP#\d+\])`)
	matches := re.FindAllStringSubmatch(pr.Title, 1)

	return len(matches) > 0

}
