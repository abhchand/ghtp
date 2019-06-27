package main

import (
	"fmt"
)

type TargetProcessAssignable struct {
	Id   int    `json:"Id"`
	Name string `json:"Name"`

	EntityState TargetProcessEntityState `json:"EntityState"`
}

type TargetProcessEntityState struct {
	Id   int    `json:"Id"`
	Name string `json:"Name"`

	NextStates struct {
		Items []TargetProcessNextState `json:"Items"`
	} `json:"NextStates"`
}

type TargetProcessNextState struct {
	Id   int    `json:"Id"`
	Name string `json:"Name"`
}

func (assignable *TargetProcessAssignable) findNextStateByName(name string) TargetProcessNextState {

	emptyNextState := TargetProcessNextState{}

	if len(name) == 0 {
		return emptyNextState
	}

	for _, nextState := range assignable.EntityState.NextStates.Items {
		if nextState.Name == name {
			return nextState
		}
	}

	return emptyNextState

}

func (assignable *TargetProcessAssignable) getCurrentEntityState() TargetProcessEntityState {

	return assignable.EntityState

}

func (assignable *TargetProcessAssignable) toString() string {

	return fmt.Sprintf("#%v %v", assignable.Id, assignable.Name)

}

func (entityState *TargetProcessEntityState) toString() string {

	return fmt.Sprintf("'%v' (id: %v)", entityState.Name, entityState.Id)

}

func (nextState *TargetProcessNextState) toString() string {

	return fmt.Sprintf("'%v' (id: %v)", nextState.Name, nextState.Id)

}
