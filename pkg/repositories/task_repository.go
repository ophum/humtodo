package repositories

import "github.com/ophum/humtodo/pkg/entities"

type TaskRepository interface {
	Find(id string) (entities.TaskEntity, error)
	FindByProjectId(projectId string) ([]entities.TaskEntity, error)
	Create(task entities.TaskEntity) (entities.TaskEntity, error)
	Update(task entities.TaskEntity) (entities.TaskEntity, error)
	AddTodo(taskId string, todo entities.TodoEntity) (entities.TaskEntity, error)
	UpdateTodo(taskId string, todo entities.TodoEntity) (entities.TaskEntity, error)
}
