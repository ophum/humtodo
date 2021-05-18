package controllers

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ophum/humtodo/pkg/services"
)

type ProjectController struct {
	projectService services.ProjectService
}

func NewProjectController(projectService services.ProjectService) *ProjectController {
	return &ProjectController{
		projectService: projectService,
	}
}

func (c *ProjectController) Index(ctx echo.Context) error {
	user := getUser(ctx)

	projects, err := c.projectService.FindJoinedAll(user.Uid)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"projects": projects,
	})
}

func (c *ProjectController) Show(ctx echo.Context) error {
	id := ctx.Param("id")
	project, tasks, err := c.projectService.FindWithTasks(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"project": project,
		"tasks":   tasks,
	})
}

// +gen-ts-entity
type CreateProjectRequest struct {
	Name string `json:"name"`
}

func (c *ProjectController) Create(ctx echo.Context) error {
	user := getUser(ctx)

	req := CreateProjectRequest{}
	ctx.Bind(&req)

	project, err := c.projectService.Create(req.Name, user.Uid)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, nil)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"project": project,
	})
}

// +gen-ts-entity
type JoinProjectRequest struct {
	UserId string `json:"user_id"`
}

func (c *ProjectController) Join(ctx echo.Context) error {
	id := ctx.Param("id")
	req := JoinProjectRequest{}
	ctx.Bind(&req)

	if err := c.projectService.Join(id, req.UserId); err != nil {
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	return ctx.JSON(http.StatusOK, nil)
}

type AuthUser struct {
	Uid  string
	Name string
	Exp  int64
}

func getUser(ctx echo.Context) AuthUser {
	user := ctx.Get(middleware.DefaultJWTConfig.ContextKey).(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	ret := AuthUser{}
	if uid, ok := claims["uid"]; ok {
		ret.Uid = uid.(string)
	}

	if name, ok := claims["name"]; ok {
		ret.Name = name.(string)
	}

	if exp, ok := claims["exp"]; ok {
		ret.Exp = int64(exp.(float64))
	}

	return ret
}
