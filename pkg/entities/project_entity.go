package entities

// +gen-ts-entity
type ProjectEntity struct {
	ID        string   `json:"_id,omitempty"`
	Name      string   `json:"name"`
	OwnerId   string   `json:"ownerId"`
	MemberIds []string `json:"members"`
}
