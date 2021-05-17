package repositories

import "github.com/ophum/humtodo/pkg/entities"

type TaskRepositoryInMemory struct {
	db []entities.TaskEntity
}

func NewTaskRepositoryInMemory() *TaskRepositoryInMemory {
	return &TaskRepositoryInMemory{
		db: []entities.TaskEntity{},
	}
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
