package entities

// +gen-ts-entity
type ProjectEntity struct {
	ID      string `json:"_id,omitempty"`
	GroupId string `json:"group_id"`
	Name    string `json:"name"`
}
