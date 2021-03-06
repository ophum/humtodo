package repositories

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/ophum/humtodo/pkg/entities"
)

type ProjectRepositoryInMemory struct {
	db []entities.ProjectEntity
}

func NewProjectRepositoryInMemory() *ProjectRepositoryInMemory {
	return &ProjectRepositoryInMemory{
		db: []entities.ProjectEntity{},
	}
}

func (r *ProjectRepositoryInMemory) Find(id string) (entities.ProjectEntity, error) {
	for _, p := range r.db {
		if p.ID == id {
			return p, nil
		}
	}

	return entities.ProjectEntity{}, fmt.Errorf("Not found")
}

func (r *ProjectRepositoryInMemory) FindAll() ([]entities.ProjectEntity, error) {
	return r.db, nil
}

func (r *ProjectRepositoryInMemory) FindJoinedAll(userId string) ([]entities.ProjectEntity, error) {
	projects := []entities.ProjectEntity{}
	for _, p := range r.db {
		for _, uid := range p.MemberIds {
			if uid == userId {
				projects = append(projects, p)
				break
			}
		}
	}
	return projects, nil
}

func (r *ProjectRepositoryInMemory) Create(project entities.ProjectEntity) (entities.ProjectEntity, error) {
	id := uuid.NewString()

	if _, err := r.Find(id); err == nil {
		return entities.ProjectEntity{}, fmt.Errorf("Already exists")
	}

	project.ID = id
	r.db = append(r.db, project)

	return r.Find(id)
}

func (r *ProjectRepositoryInMemory) Update(project entities.ProjectEntity) (entities.ProjectEntity, error) {
	for i, p := range r.db {
		if p.ID == project.ID {
			r.db[i] = project
			return r.Find(project.ID)
		}
	}
	return entities.ProjectEntity{}, fmt.Errorf("Not found")
}

func (r *ProjectRepositoryInMemory) Delete(id string) error {
	for i, p := range r.db {
		if p.ID == id {
			r.db = append(r.db[:i], r.db[i+1:]...)
			return nil
		}
	}
	return nil
}

func (r *ProjectRepositoryInMemory) IsJoinedMember(id, userId string) (bool, error) {
	project, err := r.Find(id)
	if err != nil {
		return false, err
	}

	for _, mid := range project.MemberIds {
		if mid == userId {
			return true, nil
		}
	}
	return false, nil
}
