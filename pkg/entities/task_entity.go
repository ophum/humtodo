package entities

import "time"

// +gen-ts-entity
type TaskEntity struct {
	ID          string       `json:"_id,omitempty"`
	Title       string       `json:"title"`
	Plan        int          `json:"plan"`
	Todos       []TodoEntity `json:"todos"`
	AssigneeIds []string     `json:"assignees"`
	ProjectId   string       `json:"project_id"`
}

// +gen-ts-entity
type TodoEntity struct {
	AssigneeId    string    `json:"assignee_id,omitempty"`
	StartDatetime time.Time `json:"start_datetime" ts-type:"string"`
	EndDatetime   time.Time `json:"end_datetime" ts-type:"string"`
	Description   string    `json:"description"`
}
