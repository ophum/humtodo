package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ophum/humtodo/pkg/entities"
	"github.com/ophum/humtodo/pkg/services"
)

type TaskController struct {
	projectService services.ProjectService
}

func NewTaskController(projectService services.ProjectService) *TaskController {
	return &TaskController{
		projectService: projectService,
	}
}

func (c *TaskController) Index(ctx echo.Context) error {
	return nil
}

// +gen-ts-entity
type CreateTaskRequest struct {
	Title              string   `json:"title"`
	StartDatetime      string   `json:"start_datetime"`
	EndDatetime        string   `json:"end_datetime"`
	TotalScheduledTime int      `json:"total_scheduled_time"`
	AssigneeIds        []string `json:"assignee_ids"`
}

// +gen-ts-entity
type CreateTaskResponse struct {
	Task entities.TaskEntity `json:"task" ts-import:"../entities/entities"`
}

func (c *TaskController) Create(ctx echo.Context) error {
	user := getUser(ctx)
	projId := ctx.Param("proj_id")

	if joined, err := c.projectService.IsJoined(projId, user.Uid); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	} else if !joined {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"error": "forbidden",
		})
	}

	req := CreateTaskRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	task, err := c.projectService.AddTask(projId,
		req.Title,
		req.StartDatetime,
		req.EndDatetime,
		req.TotalScheduledTime,
		req.AssigneeIds,
	)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusCreated, CreateTaskResponse{
		Task: task,
	})
}

// +gen-ts-entity
type AddTodoRequest struct {
	Title         string `json:"title"`
	AssigneeId    string `json:"assignee_id"`
	StartDatetime string `json:"start_datetime"`
	ScheduledTime int    `json:"scheduled_time"`
	Note          string `json:"note"`
}

// +gen-ts-entity
type AddTodoResponse struct {
	Task entities.TaskEntity `json:"task" ts-import:"../entities/entities"`
}

func (c *TaskController) AddTodo(ctx echo.Context) error {
	user := getUser(ctx)
	projId := ctx.Param("proj_id")
	taskId := ctx.Param("id")

	if joined, err := c.projectService.IsJoined(projId, user.Uid); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	} else if !joined {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"error": "forbidden",
		})
	}

	req := AddTodoRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	task, err := c.projectService.AddTodo(projId, taskId, req.Title, req.AssigneeId, req.Note, req.StartDatetime, req.ScheduledTime)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusCreated, AddTodoResponse{
		Task: task,
	})
}

// +gen-ts-entity
type PatchTodoRequest struct {
	TodoId        string   `json:"todo_id"`
	PatchFields   []string `json:"patch_fields"`
	Title         string   `json:"title,omitempty"`
	AssigneeId    string   `json:"assignee_id,omitempty"`
	StartDatetime string   `json:"start_datetime,omitempty"`
	ScheduledTime int      `json:"scheduled_time,omitempty"`
	ActualTime    int      `json:"actual_time,omitempty"`
	Note          string   `json:"note,omitempty"`
	IsDone        bool     `json:"is_done,omitempty"`
}

// +gen-ts-entity
type PatchTodoResponse struct {
	Task entities.TaskEntity `json:"task" ts-import:"../entities/entities"`
}

func (c *TaskController) PatchTodo(ctx echo.Context) error {
	user := getUser(ctx)
	projId := ctx.Param("proj_id")
	taskId := ctx.Param("id")

	if joined, err := c.projectService.IsJoined(projId, user.Uid); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	} else if !joined {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"error": "forbidden",
		})
	}

	req := PatchTodoRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	task, err := c.projectService.PatchTodo(
		projId,
		taskId,
		req.TodoId,
		req.PatchFields,
		req.Title,
		req.AssigneeId,
		req.StartDatetime,
		req.ScheduledTime,
		req.ActualTime,
		req.Note,
		req.IsDone,
	)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusCreated, PatchTodoResponse{
		Task: task,
	})
}
