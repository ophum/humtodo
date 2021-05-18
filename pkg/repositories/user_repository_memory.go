package repositories

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/ophum/humtodo/pkg/entities"
	"gopkg.in/yaml.v2"
)

type UserRepositoryInMemory struct {
	db []entities.UserEntity
}

func NewUserRepositoryInMemory() *UserRepositoryInMemory {
	return &UserRepositoryInMemory{
		db: []entities.UserEntity{},
	}
}

func (r *UserRepositoryInMemory) Find(id string) (entities.UserEntity, error) {
	for _, u := range r.db {
		if u.ID == id {
			return u, nil
		}
	}
	return entities.UserEntity{}, fmt.Errorf("Not found")
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

	id := uuid.NewString()
	if _, err := r.Find(id); err == nil {
		return entities.UserEntity{}, fmt.Errorf("Duplicate id")
	}

	u.ID = id
	r.db = append(r.db, u)

	return r.FindByName(u.Name)
}

func debug(data interface{}) {
	y, _ := yaml.Marshal(data)
	log.Println("\n", string(y))
}
