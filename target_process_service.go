package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func buildEntityRequest(url string) *http.Request {

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "abhchand/ghtp")

	return req

}

func entityEndpoint(entityId int) string {

	return fmt.Sprintf(
		"https://callrail.tpondemand.com/api/v1/Assignables/%v?format=json&access_token=%v",
		entityId,
		targetProcessAuthToken)

}

func targetProcessStateFor(entityId int) string {

	url := entityEndpoint(entityId)
	httpClient := http.Client{Timeout: time.Second * 2}

	// Build request
	request := buildEntityRequest(url)
	log.Debug("Querying: " + url)

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
	var entity TargetProcessEntity
	err = json.Unmarshal([]byte(body), &entity)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return entity.getState()

}
