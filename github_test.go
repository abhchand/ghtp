package main

import (
	"testing"
)

func TestExpectedTargetProcessNextStateName(t *testing.T) {

	pr := PullRequest{
		Id:      1,
		HtmlUrl: "github.com/foo",
		Title:   "My PR",
		Labels: []PullRequestLabel{
			PullRequestLabel{Name: "label1"}}}

	matchingRules := []SyncRule{
		SyncRule{IfHas: "label0", ThenSet: "state0"},
		SyncRule{IfHas: "label1", ThenSet: "state1"}}

	nonMatchingRules := []SyncRule{
		SyncRule{IfHas: "label3", ThenSet: "state0"},
		SyncRule{IfHas: "label4", ThenSet: "state1"}}

	emptyRules := []SyncRule{}

	ruleSets := [][]SyncRule{matchingRules, nonMatchingRules, emptyRules}
	expectedResult := []string{"state1", "", ""}

	for i, ruleSet := range ruleSets {

		actual := pr.expectedTargetProcessNextStateName(ruleSet)
		expected := expectedResult[i]

		if actual != expected {
			t.Error("expected", expected, "got", actual)
		}
	}

}

func TestExpectedTargetProcessNextStateNameWhenMultipleLabels(t *testing.T) {

	pr := PullRequest{
		Id:      1,
		HtmlUrl: "github.com/foo",
		Title:   "My PR",
		Labels: []PullRequestLabel{
			PullRequestLabel{Name: "label1"},
			PullRequestLabel{Name: "label0"}}}

	rules := []SyncRule{
		SyncRule{IfHas: "label0", ThenSet: "state0"},
		SyncRule{IfHas: "label1", ThenSet: "state1"}}

	actual := pr.expectedTargetProcessNextStateName(rules)
	expected := "state0"

	if actual != expected {
		t.Error("expected", expected, "got", actual)
	}
}

func TestExpectedTargetProcessNextStateNameWhenNoLabel(t *testing.T) {

	pr := PullRequest{
		Id:      1,
		HtmlUrl: "github.com/foo",
		Title:   "My PR"}

	rules := []SyncRule{
		SyncRule{IfHas: "label0", ThenSet: "state0"},
		SyncRule{IfHas: "label1", ThenSet: "state1"}}

	actual := pr.expectedTargetProcessNextStateName(rules)
	expected := ""

	if actual != expected {
		t.Error("expected", expected, "got", actual)
	}
}

func TestHasLabel(t *testing.T) {

	pr := PullRequest{
		Id:      1,
		HtmlUrl: "github.com/foo",
		Title:   "My PR",
		Labels: []PullRequestLabel{
			PullRequestLabel{Name: "label1"}}}

	testCases := map[string]bool{
		"label1": true,
		"label2": false,
		"":       false,
	}

	for labelName, expected := range testCases {

		actual := pr.hasLabel(labelName)

		if actual != expected {
			t.Error("expected", expected, "got", actual)
		}
	}

}

func TestHasLabelWhenMultipleLabels(t *testing.T) {

	pr := PullRequest{
		Id:      1,
		HtmlUrl: "github.com/foo",
		Title:   "My PR",
		Labels: []PullRequestLabel{
			PullRequestLabel{Name: "label1"},
			PullRequestLabel{Name: "label2"}}}

	testCases := map[string]bool{
		"label1": true,
		"label2": true,
		"label3": false,
		"":       false,
	}

	for labelName, expected := range testCases {

		actual := pr.hasLabel(labelName)

		if actual != expected {
			t.Error("expected", expected, "got", actual)
		}
	}

}

func TestHasLabelWhenNoLabel(t *testing.T) {

	pr := PullRequest{
		Id:      1,
		HtmlUrl: "github.com/foo",
		Title:   "My PR"}

	testCases := map[string]bool{
		"label1": false,
		"":       false,
	}

	for labelName, expected := range testCases {

		actual := pr.hasLabel(labelName)

		if actual != expected {
			t.Error("expected", expected, "got", actual)
		}
	}

}

func TestTargetShouldSync(t *testing.T) {

	testCases := map[string]bool{
		"[TP#1234] foo":     true,
		"[tp#1234] foo":     true,
		"[TP#1234]":         true,
		"foo [TP#1234] foo": false,
		"TP#1234 foo":       false,
	}

	for title, expected := range testCases {
		pr := PullRequest{
			Id:              1,
			HtmlUrl:         "github.com/foo",
			Title:           title,
			PullRequestData: map[string]string{"foo": "bar"}}

		actual := pr.shouldSync()

		if actual != expected {
			t.Error("expected", expected, "got", actual)
		}
	}

}

func TestTargetShouldSyncWhenNotTrulyPullRequest(t *testing.T) {

	testCases := map[string]bool{
		"[TP#1234] foo":     false,
		"[tp#1234] foo":     false,
		"[TP#1234]":         false,
		"foo [TP#1234] foo": false,
		"TP#1234 foo":       false,
	}

	for title, expected := range testCases {
		pr := PullRequest{
			Id:              1,
			HtmlUrl:         "github.com/foo",
			Title:           title,
			PullRequestData: map[string]string{}}

		actual := pr.shouldSync()

		if actual != expected {
			t.Error("expected", expected, "got", actual)
		}
	}

}

func TestTargetProcessAssignableId(t *testing.T) {

	testCases := map[string]int{
		"[TP#1234] foo":     1234,
		"[tp#1234] foo":     1234,
		"[TP#1234]":         1234,
		"foo [TP#1234] foo": 0,
		"TP#1234 foo":       0,
	}

	for title, expected := range testCases {
		pr := PullRequest{Id: 1, HtmlUrl: "github.com/foo", Title: title}

		actual := pr.targetProcessAssignableId()

		if actual != expected {
			t.Error("expected", expected, "got", actual)
		}
	}
}
