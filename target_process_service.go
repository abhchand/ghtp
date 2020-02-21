package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const TP_PAGE_SIZE int = 50

var targetProcessHostTemplate = "https://%v.tpondemand.com"

func createTargetProcessComment(createCommentUrl string, assignable TargetProcessAssignable, pr PullRequest) error {

	// Build Request
	payload := createTargetProcessCommentPayload(assignable, pr)
	request := createTargetProcessCommentRequestBuilder(createCommentUrl, payload)

	// Query the API
	_, err := queryTargetProcess(request)
	if err != nil {
		return err
	}

	log.Debugf("[%v] Created TargetProcess Comment on #%v",
		pr.toString(),
		assignable.Id)

	return nil

}

func createTargetProcessCommentPayload(assignable TargetProcessAssignable, pr PullRequest) string {

	comment := fmt.Sprintf(
		"Moved to '%v' based on [%v](%v)",
		assignable.EntityState.Name,
		pr.toString(),
		pr.HtmlUrl)

	return fmt.Sprintf(
		"{ General: { Id: %v }, Description: \"<!--markdown-->%v\" }",
		assignable.Id,
		comment)

}

func createTargetProcessCommentRequestBuilder(url string, payload string) *http.Request {

	body := bytes.NewBuffer([]byte(payload))

	req, err := http.NewRequest(http.MethodPost, url, body)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "abhchand/ghtp")
	req.Header.Set("Content-Type", "application/json")

	return req

}

func createTargetProcessCommentUrl() string {

	return fmt.Sprintf(
		"%v/api/v1/Comments?access_token=%v",
		targetProcessHost(),
		targetProcessAuthToken)

}

func updateTargetProcessEntityState(updateEntityStateUrl string, pr PullRequest, assignable TargetProcessAssignable, nextState TargetProcessNextState) TargetProcessAssignable {

	// Build Reques
	payload := updateTargetProcessEntityStatePayload(nextState)
	request := updateTargetProcessEntityStateRequestBuilder(updateEntityStateUrl, payload)

	// Query API
	_, err := queryTargetProcess(request)
	if err != nil {
		log.Error(err.Error())
		return TargetProcessAssignable{}
	}

	log.Infof("[%v] Updated TargetProcess #%v to state '%v' ☑️",
		pr.toString(),
		assignable.Id,
		nextState.toString())

	// Construct an updated assignable
	assignable.EntityState = TargetProcessEntityState{
		Id: nextState.Id, Name: nextState.Name}

	return assignable

}

func updateTargetProcessEntityStatePayload(nextState TargetProcessNextState) string {

	return fmt.Sprintf("{ EntityState:{Id:%v} }", nextState.Id)

}

func updateTargetProcessEntityStateRequestBuilder(url string, payload string) *http.Request {

	body := bytes.NewBuffer([]byte(payload))

	req, err := http.NewRequest(http.MethodPost, url, body)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "abhchand/ghtp")
	req.Header.Set("Content-Type", "application/json")

	return req

}

func updateTargetProcessEntityStateUrl(assignable TargetProcessAssignable) string {

	return fmt.Sprintf(
		"%v/api/v1/Assignables/%v?"+
			"resultFormat=json&resultInclude=[Id]&access_token=%v",
		targetProcessHost(),
		assignable.Id,
		targetProcessAuthToken)

}

func findTargetProcessAssignable(id int) TargetProcessAssignable {

	// Build Request
	url := findTargetProcessAssignableUrl(id)
	request := findTargetProcessAssignableRequestBuilder(url)

	// Query API
	response, err := queryTargetProcess(request)
	if err != nil {
		log.Error(err.Error())
		return TargetProcessAssignable{}
	}

	// Read response body
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	// Load Response
	var assignable TargetProcessAssignable
	err = json.Unmarshal([]byte(responseBody), &assignable)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return assignable

}

func findTargetProcessAssignableRequestBuilder(url string) *http.Request {

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "abhchand/ghtp")

	return req

}

func findTargetProcessAssignableUrl(assignableId int) string {

	return fmt.Sprintf(
		"%v/api/v1/Assignables/%v?"+
			"include=[Name,EntityState[Name,NextStates[Id,Name]]]&"+
			"format=json&access_token=%v",
		targetProcessHost(),
		assignableId,
		targetProcessAuthToken)

}

func queryTargetProcess(request *http.Request) (*http.Response, error) {
	log.Debugf("[%v] %v (%v)", request.Method, request.URL.String(), request.Body)

	httpClient := http.Client{Timeout: time.Second * 2}

	// Query API
	response, err := httpClient.Do(request)
	if err != nil {
		return response, err
	}

	// Handle bad HTTP response
	if response.StatusCode < 200 || response.StatusCode > 299 {
		msg := fmt.Sprintf("Error querying API (response: %s)", response.Status)
		return response, errors.New(msg)
	}

	return response, nil

}

func targetProcessHost() string {

	return fmt.Sprintf(targetProcessHostTemplate, targetProcessDomain)

}
