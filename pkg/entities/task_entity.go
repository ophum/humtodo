package entities

// +gen-ts-entity
type TaskEntity struct {
	ID                 string       `json:"_id,omitempty"`
	Title              string       `json:"title"`
	TotalScheduledTime int          `json:"total_scheduled_time"`
	Todos              []TodoEntity `json:"todos"`
	AssigneeIds        []string     `json:"assignee_ids"`
	ProjectId          string       `json:"project_id"`
}

// +gen-ts-entity
type TodoEntity struct {
	AssigneeId    string `json:"assignee_id,omitempty"`
	StartDatetime string `json:"start_datetime"`
	ScheduledTime int    `json:"scheduled_time"`
	ActualTime    int    `json:"actual_time"`
	Description   string `json:"description"`
}
