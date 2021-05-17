package entities

type UserEntity struct {
	ID       string `json:"_id,omitempty"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
