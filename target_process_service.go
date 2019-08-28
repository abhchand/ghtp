package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const TP_PAGE_SIZE int = 50

var targetProcessHostTemplate = "https://%v.tpondemand.com"

func assignableEndpoint(assignableId int) string {

	return fmt.Sprintf(
		"%v/api/v1/Assignables/%v?"+
			"include=[Name,EntityState[Name,NextStates[Id,Name]]]&"+
			"format=json&access_token=%v",
		targetProcessHost(),
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

func buildCommentRequest(url string, payload string) *http.Request {

	body := bytes.NewBuffer([]byte(payload))

	req, err := http.NewRequest(http.MethodPost, url, body)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "abhchand/ghtp")
	req.Header.Set("Content-Type", "application/json")

	return req

}

func buildUpdateEntityStateRequest(url string, payload string) *http.Request {

	body := bytes.NewBuffer([]byte(payload))

	req, err := http.NewRequest(http.MethodPost, url, body)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "abhchand/ghtp")
	req.Header.Set("Content-Type", "application/json")

	return req

}

func commentEndpoint() string {

	return fmt.Sprintf(
		"%v/api/v1/Comments?access_token=%v",
		targetProcessHost(),
		targetProcessAuthToken)

}

func createCommentPayload(assignable TargetProcessAssignable, pr PullRequest) string {

	comment := fmt.Sprintf(
		"Moved to '%v' based on labels in [%v](%v)",
		assignable.EntityState.Name,
		pr.toString(),
		pr.HtmlUrl)

	return fmt.Sprintf(
		"{ General: { Id: %v }, Description: \"<!--markdown-->%v\" }",
		assignable.Id,
		comment)

}

func createTargetProcessComment(assignable TargetProcessAssignable, pr PullRequest) {

	// Build Request
	url := commentEndpoint()
	payload := createCommentPayload(assignable, pr)
	request := buildCommentRequest(url, payload)

	queryTargetProcess(request)

	// `queryTargetProcess()` exits or panics if there's an error, so assume
	// everything is successful at this point
	log.Debugf("[%v] Created TargetProcess Comment on #%v",
		pr.toString(),
		assignable.Id)

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

	log.Debugf("[%v] %v", request.Method, request.URL.String())

	httpClient := http.Client{Timeout: time.Second * 2}

	// Query API
	response, err := httpClient.Do(request)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer response.Body.Close()

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

func targetProcessHost() string {

	return fmt.Sprintf(targetProcessHostTemplate, targetProcessDomain)

}

func updateEntityStateEndpoint(assignable TargetProcessAssignable) string {

	return fmt.Sprintf(
		"%v/api/v1/Assignables/%v?"+
			"resultFormat=json&resultInclude=[Id]&access_token=%v",
		targetProcessHost(),
		assignable.Id,
		targetProcessAuthToken)

}

func updateEntityStatePayload(nextState TargetProcessNextState) string {

	return fmt.Sprintf("{ EntityState:{Id:%v} }", nextState.Id)

}

func updateTargetProcessEntityState(pr PullRequest, assignable TargetProcessAssignable, nextState TargetProcessNextState) TargetProcessAssignable {

	// Build Request
	url := updateEntityStateEndpoint(assignable)
	payload := updateEntityStatePayload(nextState)
	request := buildUpdateEntityStateRequest(url, payload)

	queryTargetProcess(request)

	// `queryTargetProcess()` exits or panics if there's an error, so assume
	// everything is successful at this point
	log.Infof("[%v] Updated TargetProcess #%v to state '%v' ☑️",
		pr.toString(),
		assignable.Id,
		nextState.toString())

	// Construct an updated assignable
	assignable.EntityState = TargetProcessEntityState{
		Id: nextState.Id, Name: nextState.Name}

	return assignable

}
