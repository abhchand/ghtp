package main

type TargetProcessEntity struct {
	ResourceType string `json:"ResourceType"`
	ID           int    `json:"Id"`
	Name         string `json:"Name"`

	EntityState struct {
		State string `json:"Name"`
	} `json:"EntityState"`
}

func (entity *TargetProcessEntity) getState() string {

	return entity.EntityState.State

}
