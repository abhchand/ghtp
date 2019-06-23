package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"time"
)

func buildIssueRequest(url string) *http.Request {

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "abhchand/ghtp")
	req.Header.Set("Authorization", "token "+githubAuthToken)

	return req

}

func findEligiblePullRequests() PullRequestList {

	httpClient := http.Client{Timeout: time.Second * 2}

	var prList PullRequestList
	url := issuesEndpoint()

	// Loop through each page
	for {

		// Build request
		request := buildIssueRequest(url)
		log.Info("Querying: " + url)

		// Query API
		response, err := httpClient.Do(request)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}

		// Handle bad HTTP response
		log.Debugf("Response Status: %s", response.Status)
		if response.StatusCode < 200 || response.StatusCode > 299 {
			log.Fatal("Error querying API. Exiting")
			os.Exit(1)
		}

		// Parse response
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}

		// Load Response
		var prs PullRequestList
		err = json.Unmarshal([]byte(body), &prs)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}

		// Consider only the Pull Requests marked for sync
		for _, pr := range prs {
			if pr.shouldSync() {
				prList = append(prList, pr)
			}
		}

		// Check if we should continue (whether a next page exists)
		nextUrl := parseNextUrl(response.Header["Link"])
		log.Debugf("Next Page: %s", nextUrl)

		if len(nextUrl) > 0 {
			url = nextUrl
		} else {
			break
		}

	}

	return prList

}

func issuesEndpoint() string {

	return fmt.Sprintf(
		"https://api.github.com/repos/%s/%s/issues?state=open&direction=desc",
		githubOrganization,
		githubRepository)

}

func parseNextUrl(linkHeader []string) string {

	// Github responsds with a link header that contains meta information
	// about the next page and the last page
	//
	//     <https://api.github.com/resource?page=2>; rel="next", \n
	//       <https://api.github.com/resource?page=5>; rel="last"
	//
	// 1. Link header may not always be present. If there's not one, there is
	//    no next page so return nil
	//
	// 2. Github conveniently returns the full URL required to query the next
	//    page, including any query params you originally included. So we
	//    extract that URL if it exists and rely on it blindly.
	//
	// 3. Golang's http returns the link header as a slice of strings. Not sure
	//    if it would ever have more than one element, but we only care about
	//    the first element/string

	if len(linkHeader) == 0 {
		return ""
	}

	re := regexp.MustCompile(`(http[^,\s]*)\>; rel=\"next\"`)
	matches := re.FindAllStringSubmatch(linkHeader[0], 1)

	if len(matches) > 0 {
		return matches[0][1]
	}

	return ""

}
