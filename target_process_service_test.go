package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateTargetProcessComment(t *testing.T) {
	log = initializeLogger()

	// Setup Data
	entityState := TargetProcessEntityState{Id: 99, Name: "Development"}
	assignable := TargetProcessAssignable{Id: 21, Name: "My Cool Story", EntityState: entityState}
	pr := PullRequest{Id: 1, HtmlUrl: "github.com/foo", Title: "My PR"}

	t.Run("Success", func(t *testing.T) {

		// Setup Server
		handlerFunc := func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(200)
			res.Write([]byte("Fake Body"))
		}
		testServer := httptest.NewServer(http.HandlerFunc(handlerFunc))
		defer testServer.Close()

		err := createTargetProcessComment(testServer.URL, assignable, pr)

		if err != nil {
			t.Error("expected", nil, "got", err)
		}

	})

	t.Run("HTTP Response is not 2XX", func(t *testing.T) {

		// Setup Server
		handlerFunc := func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(500)
			res.Write([]byte("Fake Body"))
		}
		testServer := httptest.NewServer(http.HandlerFunc(handlerFunc))
		defer testServer.Close()

		err := createTargetProcessComment(testServer.URL, assignable, pr)

		if err == nil {
			t.Error("expected", "Error querying API", "got", err)
		}

	})

}

func TestCreateTargetProcessCommentPayload(t *testing.T) {
	log = initializeLogger()

	// Setup Data
	// Setup Data
	entityState := TargetProcessEntityState{Id: 99, Name: "Development"}
	assignable := TargetProcessAssignable{Id: 21, Name: "My Cool Story", EntityState: entityState}
	pr := PullRequest{Id: 1, HtmlUrl: "github.com/foo", Title: "My PR"}

	t.Run("Success", func(t *testing.T) {

		actual := createTargetProcessCommentPayload(assignable, pr)
		expected := "{ General: { Id: 21 }, Description: \"" +
			"<!--markdown-->Moved to 'Development' based on labels in [/#1](github.com/foo)\" }"

		if actual != expected {
			t.Error("expected", expected, "got", actual)
		}

	})

}
