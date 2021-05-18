package entities

// +gen-ts-entity
type TaskEntity struct {
	ID          string       `json:"_id,omitempty"`
	Title       string       `json:"title"`
	Plan        int          `json:"plan"`
	Todos       []TodoEntity `json:"todos"`
	AssigneeIds []string     `json:"assignee_ids"`
	ProjectId   string       `json:"project_id"`
}

// +gen-ts-entity
type TodoEntity struct {
	AssigneeId    string `json:"assignee_id,omitempty"`
	StartDatetime string `json:"start_datetime"`
	EndDatetime   string `json:"end_datetime"`
	Description   string `json:"description"`
}
