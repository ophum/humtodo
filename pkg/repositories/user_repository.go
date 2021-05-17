package repositories

import (
	"fmt"
	"log"

	"github.com/ophum/humtodo/pkg/entities"
	"gopkg.in/yaml.v2"
)

type UserRepository interface {
	FindByName(name string) (entities.UserEntity, error)
	Create(u entities.UserEntity) (entities.UserEntity, error)
}

type UserRepositoryInMemory struct {
	db []entities.UserEntity
}

func NewUserRepositoryInMemory() *UserRepositoryInMemory {
	return &UserRepositoryInMemory{
		db: []entities.UserEntity{},
	}
}

func (r *UserRepositoryInMemory) FindByName(name string) (entities.UserEntity, error) {
	for _, u := range r.db {
		if u.Name == name {
			return u, nil
		}
	}
	return entities.UserEntity{}, fmt.Errorf("Not found")
}

func (r *UserRepositoryInMemory) Create(u entities.UserEntity) (entities.UserEntity, error) {
	if _, err := r.FindByName(u.Name); err == nil {
		return entities.UserEntity{}, fmt.Errorf("Already exists")
	}

	r.db = append(r.db, u)

	debug(r.db)
	return r.FindByName(u.Name)
}

func debug(data interface{}) {
	y, _ := yaml.Marshal(data)
	log.Println("\n", string(y))
}
