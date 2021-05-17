package entities

// +gen-ts-entity
type ProjectEntity struct {
	ID      string       `json:"_id,omitempty"`
	Name    string       `json:"name"`
	Members []UserEntity `json:"members"`
}