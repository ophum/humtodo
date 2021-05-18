package services

import (
	"github.com/ophum/humtodo/pkg/entities"
	"github.com/ophum/humtodo/pkg/repositories"
)

type ProjectService struct {
	projectRepo repositories.ProjectRepository
	taskRepo    repositories.TaskRepository
}

func NewProjectService(projectRepo repositories.ProjectRepository, taskRepo repositories.TaskRepository) *ProjectService {
	return &ProjectService{
		projectRepo: projectRepo,
		taskRepo:    taskRepo,
	}
}

func (s *ProjectService) FindAll() ([]entities.ProjectEntity, error) {
	return s.projectRepo.FindAll()
}

func (s *ProjectService) FindWithTasks(id string) (entities.ProjectEntity, []entities.TaskEntity, error) {
	project, err := s.projectRepo.Find(id)
	if err != nil {
		return entities.ProjectEntity{},
			[]entities.TaskEntity{},
			err
	}

	tasks, err := s.taskRepo.FindByProjectId(project.ID)
	if err != nil {
		return entities.ProjectEntity{},
			[]entities.TaskEntity{},
			err
	}
	return project, tasks, nil
}

func (s *ProjectService) FindJoinedAll(userId string) ([]entities.ProjectEntity, error) {
	return s.projectRepo.FindJoinedAll(userId)
}

func (s *ProjectService) Create(name, uid string) (entities.ProjectEntity, error) {
	return s.projectRepo.Create(entities.ProjectEntity{
		Name:    name,
		OwnerId: uid,
		MemberIds: []string{
			uid,
		},
	})
}

func (s *ProjectService) IsJoined(id, userId string) (bool, error) {
	return s.projectRepo.IsJoinedMember(id, userId)
}

func (s *ProjectService) Join(id, userId string) error {
	isJoined, err := s.projectRepo.IsJoinedMember(id, userId)
	if err != nil {
		return err
	}

	if isJoined {
		return nil
	}

	project, err := s.projectRepo.Find(id)
	if err != nil {
		return err
	}

	project.MemberIds = append(project.MemberIds, userId)

	_, err = s.projectRepo.Update(project)
	return err
}

func (s *ProjectService) AddTask(projectId, title string, plan int, assigneeIds []string) (entities.TaskEntity, error) {
	return s.taskRepo.Create(entities.TaskEntity{
		Title:       title,
		Plan:        plan,
		AssigneeIds: assigneeIds,
		ProjectId:   projectId,
	})
}
