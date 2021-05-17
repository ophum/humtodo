package entities

// +gen-ts-entity
type GroupEntity struct {
	ID          string   `json:"_id,omitempty"`
	Name        string   `json:"name"`
	OwnerUserId string   `json:"owner_user_id"`
	Members     []string `json:"members"`
}
