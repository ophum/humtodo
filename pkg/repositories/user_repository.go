package repositories

import (
	"github.com/ophum/humtodo/pkg/entities"
)

type UserRepository interface {
	FindByName(name string) (entities.UserEntity, error)
	Create(u entities.UserEntity) (entities.UserEntity, error)
}
