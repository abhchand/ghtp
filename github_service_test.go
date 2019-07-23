package main

import (
	"testing"
)

func TestParseNextUrl(t *testing.T) {

	testCases := map[string]string{
		"<https://api.github.com/resource?page=2>; rel=\"next\", \n<https://api.github.com/resource?page=5>; rel=\"last\"": "https://api.github.com/resource?page=2",
		"<https://api.github.com/resource?page=2>; rel=\"last\", \n<https://api.github.com/resource?page=5>; rel=\"next\"": "https://api.github.com/resource?page=5",
		"<https://api.github.com/resource?page=2>; rel=\"next\"":                                                           "https://api.github.com/resource?page=2",
		"<https://api.github.com/resource?page=2>; rel=\"last\"":                                                           "",
		"": "",
	}

	for input, expected := range testCases {

		actual := parseNextUrl([]string{input})

		if actual != expected {
			t.Error("expected", expected, "got", actual)
		}
	}

}
