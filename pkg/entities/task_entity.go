package entities

import "time"

// +gen-ts-entity
type TaskEntity struct {
	ID        string       `json:"_id,omitempty"`
	Title     string       `json:"title"`
	Plan      int          `json:"plan"`
	Todos     []TodoEntity `json:"todos"`
	Assignees []UserEntity `json:"assignees"`
}

// +gen-ts-entity
type TodoEntity struct {
	StartDatetime time.Time `json:"start_datetime" ts-type:"string"`
	EndDatetime   time.Time `json:"end_datetime" ts-type:"string"`
	Description   string    `json:"description"`
}
