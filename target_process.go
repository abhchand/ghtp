package main

type TargetProcessEntity struct {
	ResourceType string `json:"ResourceType"`
	ID           int    `json:"Id"`
	Name         string `json:"Name"`

	State TargetProcessEntityState `json:"EntityState"`
}

type TargetProcessEntityState struct {
	Id   int    `json:"Id"`
	Name string `json:"Name"`
}

type TargetProcessEntityStateList []TargetProcessEntityState

type TargetProcessEntityStateApiResponse struct {
	Next  string                       `json:"Next"`
	Items TargetProcessEntityStateList `json:"Items"`
}

func (entity *TargetProcessEntity) getState() string {

	return entity.State.Name

}

func (entities *TargetProcessEntityStateList) findByName(name string) TargetProcessEntityState {

	emptyEntity := TargetProcessEntityState{}

	if len(name) == 0 {
		return emptyEntity
	}

	for _, entity := range *entities {
		if entity.Name == name {
			return entity
		}
	}

	return emptyEntity

}
