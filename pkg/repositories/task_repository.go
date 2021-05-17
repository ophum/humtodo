package repositories

import "github.com/ophum/humtodo/pkg/entities"

type TaskRepository interface {
	Find(id string) (entities.TaskEntity, error)
	FindByProjectId(projectId string) ([]entities.TaskEntity, error)
	Create(task entities.TaskEntity) (entities.TaskEntity, error)
}
