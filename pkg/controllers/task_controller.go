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
	Title       string   `json:"title"`
	Plan        int      `json:"plan"`
	AssigneeIds []string `json:"assignee_ids"`
}

// +gen-ts-entity
type CreateTaskResponse struct {
	Task entities.TaskEntity `json:"task"`
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

	task, err := c.projectService.AddTask(projId, req.Title, req.Plan, req.AssigneeIds)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusCreated, CreateTaskResponse{
		Task: task,
	})
}
