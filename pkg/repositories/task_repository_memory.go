package repositories

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/ophum/humtodo/pkg/entities"
)

type TaskRepositoryInMemory struct {
	db []entities.TaskEntity
}

func NewTaskRepositoryInMemory() *TaskRepositoryInMemory {
	return &TaskRepositoryInMemory{
		db: []entities.TaskEntity{},
	}
}

func (r *TaskRepositoryInMemory) Find(id string) (entities.TaskEntity, error) {
	for _, t := range r.db {
		if t.ID == id {
			return t, nil
		}
	}
	return entities.TaskEntity{}, fmt.Errorf("Not found")
}

func (r *TaskRepositoryInMemory) FindByProjectId(projectId string) ([]entities.TaskEntity, error) {
	tasks := []entities.TaskEntity{}
	for _, t := range r.db {
		if t.ProjectId == projectId {
			tasks = append(tasks, t)
		}
	}

	return tasks, nil
}

func (r *TaskRepositoryInMemory) Create(task entities.TaskEntity) (entities.TaskEntity, error) {
	id := uuid.NewString()

	if _, err := r.Find(id); err == nil {
		return entities.TaskEntity{}, fmt.Errorf("Already exists")
	}

	task.ID = id
	r.db = append(r.db, task)

	return r.Find(id)
}
