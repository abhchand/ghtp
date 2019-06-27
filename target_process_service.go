package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const TP_PAGE_SIZE int = 50

var targetProcessBase = "https://callrail.tpondemand.com"

func assignableEndpoint(assignableId int) string {

	return fmt.Sprintf(
		"%v/api/v1/Assignables/%v?"+
			"include=[Name,EntityState[Name,NextStates[Id,Name]]]&"+
			"format=json&access_token=%v",
		targetProcessBase,
		assignableId,
		targetProcessAuthToken)

}

func buildAssignableRequest(url string) *http.Request {

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "abhchand/ghtp")

	return req

}

func findTargetProcessAssignableById(id int) TargetProcessAssignable {

	// Build Request
	url := assignableEndpoint(id)
	request := buildAssignableRequest(url)

	responseBody := queryTargetProcess(request)

	// Load Response
	var assignable TargetProcessAssignable
	err := json.Unmarshal([]byte(responseBody), &assignable)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return assignable

}

func queryTargetProcess(request *http.Request) []byte {

	log.Debug("Querying: " + request.URL.String())

	httpClient := http.Client{Timeout: time.Second * 2}

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

	return body

}
