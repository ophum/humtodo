package services

import (
	"fmt"
	"sort"
	"time"

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
	for i := range tasks {
		sort.Slice(tasks[i].Todos, func(j, k int) bool {
			a, _ := time.Parse("2006-01-06T15:04", tasks[i].Todos[j].StartDatetime)
			b, _ := time.Parse("2006-01-06T15:04", tasks[i].Todos[k].StartDatetime)
			return a.Unix() < b.Unix()
		})
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

func (s *ProjectService) AddTask(projectId, title, startDatetime, endDatetime string, totalScheduledTime int, assigneeIds []string) (entities.TaskEntity, error) {
	return s.taskRepo.Create(entities.TaskEntity{
		Title:              title,
		StartDatetime:      startDatetime,
		EndDatetime:        endDatetime,
		TotalScheduledTime: totalScheduledTime,
		AssigneeIds:        assigneeIds,
		ProjectId:          projectId,
	})
}

func (s *ProjectService) AddTodo(projId, taskId, title, assigneeId, note, startDatetime string, scheduledTime int) (entities.TaskEntity, error) {
	task, err := s.taskRepo.Find(taskId)
	if err != nil {
		return entities.TaskEntity{}, err
	}

	if task.ProjectId != projId {
		return entities.TaskEntity{}, fmt.Errorf("Not found")
	}

	return s.taskRepo.AddTodo(taskId, entities.TodoEntity{
		Title:         title,
		AssigneeId:    assigneeId,
		StartDatetime: startDatetime,
		ScheduledTime: scheduledTime,
		ActualTime:    0,
		Note:          note,
		IsDone:        false,
	})
}

func (s *ProjectService) UpdateIsDoneTodo(projId, taskId, todoId string, isDone bool) (entities.TaskEntity, error) {
	task, err := s.taskRepo.Find(taskId)
	if err != nil {
		return entities.TaskEntity{}, err
	}

	if task.ProjectId != projId {
		return entities.TaskEntity{}, fmt.Errorf("Not found")
	}

	for _, todo := range task.Todos {
		if todo.ID == todoId {
			todo.IsDone = isDone
			return s.taskRepo.UpdateTodo(taskId, todo)
		}
	}
	return entities.TaskEntity{}, fmt.Errorf("Not found")

}
