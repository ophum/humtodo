package repositories

import "github.com/ophum/humtodo/pkg/entities"

type TaskRepository interface {
	FindByProjectId(projectId string) ([]entities.TaskEntity, error)
}
