package main

import (
	"testing"
)

func TestFindNextStateByName(t *testing.T) {

	assignable := TargetProcessAssignable{
		Id:   21,
		Name: "My Cool Story",
		EntityState: TargetProcessEntityState{
			Id:   99,
			Name: "Development",
		},
	}

	ns1 := TargetProcessNextState{Id: 1, Name: "Code Review"}
	ns2 := TargetProcessNextState{Id: 2, Name: "Ready To Ship"}
	empty := TargetProcessNextState{}

	assignable.EntityState.NextStates.Items = []TargetProcessNextState{ns1, ns2}

	testCases := map[string]TargetProcessNextState{
		"Code Review":   ns1,
		"Ready To Ship": ns2,
		"":              empty,
	}

	for input, expected := range testCases {

		actual := assignable.findNextStateByName(input)

		if actual != expected {
			t.Error("expected", expected, "got", actual)
		}
	}

}
