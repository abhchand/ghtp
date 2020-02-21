package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type PullRequest struct {
	Id       int    `json:"number"`
	HtmlUrl  string `json:"html_url"`
	Title    string `json:"title"`
	MergedAt string `json:"merged_at"`

	Labels          []PullRequestLabel `json:"labels"`
	PullRequestData map[string]string  `json:"pull_request"`
}

type PullRequestLabel struct {
	Name string `json:"name"`
}

type PullRequestList []PullRequest

// Returns the expected entity state of the associated TargetProcess story by
// applying the rules in the config YML file.
//
// Rules are applied in the defined order and the first succesful match is
// returned.
//
// The special keyword `:pr_merged:` can be used to identify the state when
// the PR is merged (different from closed).
//
// All labels and states are case sensitive
//
// Example:
//
// Assume the following config defined in the YAML file
//
//     sync:
//       - if_has: :pr_merged:
//         then_set: Shipped
//       - if_has: chennai
//         then_set: Development
//       - if_has: bangalore
//         then_set: Shipped
//       - if_has: mumbai
//         then_set: Code Review
//
// If an open `PullRequest` had labels ["bangalore", "mumbai"], the expected
// TargetProcess state returned would be "Shipped"
//
func (pr *PullRequest) expectedTargetProcessNextStateName(rules []SyncRule) string {

	for _, rule := range rules {
		label := rule.IfHas
		state := rule.ThenSet

		if label == ":pr_merged:" && pr.isMerged() {
			return state
		} else if pr.hasLabel(label) {
			return state
		}
	}

	return ""

}

func (pr *PullRequest) hasLabel(label string) bool {

	for _, prLabel := range pr.Labels {
		if prLabel.Name == label {
			return true
		}
	}

	return false

}

func (pr *PullRequest) isMerged() bool {
	return len(pr.MergedAt) > 0
}

// Check whether this Pull Request is eligible to be included
// in the sync process by looking at the following conditions:
//
//   - If PR is truly a Pull Request and not an Issue (see
//     `isTrulyPullRequest() function)
//   - If PR Title starts with '[TP#1234]', case insensitive.
//
func (pr *PullRequest) shouldSync() bool {

	re := regexp.MustCompile(`(?i)^(\[TP#\d+\])`)
	matches := re.FindAllStringSubmatch(pr.Title, 1)

	return len(matches) > 0

}

func (pr *PullRequest) targetProcessAssignableId() int {

	re := regexp.MustCompile(`(?i)^\[TP#(\d+)\]`)
	matches := re.FindAllStringSubmatch(pr.Title, 1)

	if matches == nil {
		return 0
	}

	id, err := strconv.Atoi(matches[0][1])
	if err != nil {
		panic(err)
	}

	return id

}

func (pr *PullRequest) toString() string {

	return fmt.Sprintf("%v/%v#%v", githubOrganization, githubRepository, pr.Id)

}
