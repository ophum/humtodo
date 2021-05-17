package controllers

import (
	"net/http"

	"github.com/labstack/echo"
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
	Title string `json:"title"`
}

func (c *TaskController) Create(ctx echo.Context) error {
	projId := ctx.Param("proj_id")
	req := CreateTaskRequest{}
	ctx.Bind(&req)

	task, err := c.projectService.AddTask(projId, req.Title)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"task": task,
	})
}
