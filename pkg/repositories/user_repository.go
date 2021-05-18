package repositories

import (
	"github.com/ophum/humtodo/pkg/entities"
)

type UserRepository interface {
	Find(id string) (entities.UserEntity, error)
	FindByName(name string) (entities.UserEntity, error)
	Create(u entities.UserEntity) (entities.UserEntity, error)
}
