package main

import (
	"regexp"
)

type PullRequest struct {
	Id      int    `json:"number"`
	HtmlUrl string `json:"html_url"`
	Title   string `json:"title"`

	Labels          []PullRequestLabel  `json:"labels"`
	PullRequestData map[string]string `json:"pull_request"`
}

type PullRequestList []PullRequest

// NOTE: Github API treats all pull requests as Issues. That is, all pull
// requests are issues but not all issues are pull requests.
//
// In fact, the underlying API endpoint we are querying is for issues, but
// we are modeling this data structure as a Pull Request for convenience
// sake.
//
// To dinstiguish which issues are in fact truly pull requests, Github
// recommends looking for the presence of a `pull_request` key, which
// is what this function looks for
func (pr *PullRequest) IsTrulyPullRequest() bool {

	return len(pr.PullRequestData) > 0

}

// Check whether this Pull Request is eligible to be included
// in the sync process by looking at the following conditions:
//
//   - If PR is truly a Pull Request and not an Issue (see
//     `IsTrulyPullRequest() function)
//   - If PR Title starts with '[TP#1234]', case insensitive.
//
func (pr *PullRequest) ShouldSync() bool {

	if !pr.IsTrulyPullRequest() {
		return false
	}

	re := regexp.MustCompile(`^(\[TP#\d+\])`)
	matches := re.FindAllStringSubmatch(pr.Title, 1)

	return len(matches) > 0

}
