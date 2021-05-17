package controllers

import "github.com/labstack/echo"

type TaskController struct {
}

func NewTaskController() *TaskController {
	return &TaskController{}
}

func (c *TaskController) Index(ctx echo.Context) error {
	return nil
}
