package entities

// +gen-ts-entity
type UserEntity struct {
	ID       string `json:"_id,omitempty"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
