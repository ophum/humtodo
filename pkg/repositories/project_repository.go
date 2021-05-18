package repositories

import (
	"github.com/ophum/humtodo/pkg/entities"
)

type ProjectRepository interface {
	Find(id string) (entities.ProjectEntity, error)
	FindAll() ([]entities.ProjectEntity, error)
	FindJoinedAll(userId string) ([]entities.ProjectEntity, error)
	Create(project entities.ProjectEntity) (entities.ProjectEntity, error)
	Update(project entities.ProjectEntity) (entities.ProjectEntity, error)
	Delete(id string) error

	IsJoinedMember(id, userId string) (bool, error)
}
