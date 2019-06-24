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
var allTargetProcessEntityStates TargetProcessEntityStateList

func buildEntityRequest(url string) *http.Request {

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "abhchand/ghtp")

	return req

}

func buildEntityStateRequest(url string) *http.Request {

	return buildEntityRequest(url)

}

func entityEndpoint(entityId int) string {

	return fmt.Sprintf(
		"%v/api/v1/Assignables/%v?format=json&access_token=%v",
		targetProcessBase,
		entityId,
		targetProcessAuthToken)

}

func entityStateEndpoint(page int) string {

	return fmt.Sprintf(
		"%v/api/v1/EntityStates?format=json&take=%v&skip=%v&access_token=%v",
		targetProcessBase,
		TP_PAGE_SIZE,
		(page-1)*TP_PAGE_SIZE,
		targetProcessAuthToken)

}

func fetchAllTargetProcessEntityStates() TargetProcessEntityStateList {

	var allStates TargetProcessEntityStateList
	page := 1

	for {

		// Build Request
		url := entityStateEndpoint(page)
		request := buildEntityStateRequest(url)

		responseBody := queryTargetProcess(request)

		// Load Response
		var tpApiResponse TargetProcessEntityStateApiResponse
		err := json.Unmarshal([]byte(responseBody), &tpApiResponse)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}

		// Append to list
		allStates = append(allStates, tpApiResponse.Items...)

		// Check if we should continue (whether a next page exists)

		log.Debugf("Next Page: %s", tpApiResponse.Next)
		if tpApiResponse.Next == "" {
			break
		}

		page++

	}

	allTargetProcessEntityStates = allStates
	log.Debugf("Found EntityStates: %v", allTargetProcessEntityStates)

	return allTargetProcessEntityStates

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

func targetProcessStateFor(entityId int) string {

	// Build Request
	url := entityEndpoint(entityId)
	request := buildEntityRequest(url)

	responseBody := queryTargetProcess(request)

	// Load Response
	var entity TargetProcessEntity
	err := json.Unmarshal([]byte(responseBody), &entity)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return entity.getState()

}
