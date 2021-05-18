package routes

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ophum/humtodo/pkg/controllers"
	"github.com/ophum/humtodo/pkg/repositories"
	"github.com/ophum/humtodo/pkg/services"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())

	userRepo := repositories.NewUserRepositoryInMemory()
	projectRepo := repositories.NewProjectRepositoryInMemory()
	taskRepo := repositories.NewTaskRepositoryInMemory()

	api := e.Group("/api")
	{
		authRoutes(api, userRepo)
		projectRoutes(api, projectRepo, taskRepo)
	}

	return e
}

func authRoutes(e *echo.Group, userRepo repositories.UserRepository) {
	authService := services.NewAuthService([]byte("test"), &userRepo)
	authController := controllers.NewAuthController(*authService)
	g := e.Group("/auth")
	{
		g.POST("/sign-in", authController.SignIn)
		g.POST("/sign-up", authController.SignUp)

		auth := g.Group("")
		auth.Use(middleware.JWT([]byte("test")))
		{
			auth.POST("/verify", authController.Verify)
		}
	}
}

func projectRoutes(e *echo.Group, projectRepo repositories.ProjectRepository, taskRepo repositories.TaskRepository) {
	projectService := services.NewProjectService(projectRepo, taskRepo)
	projController := controllers.NewProjectController(*projectService)
	g := e.Group("/projects")
	g.Use(middleware.JWT([]byte("test")))
	{
		g.GET("", projController.Index)
		g.GET("/:id", projController.Show)
		g.POST("", projController.Create)
		g.PATCH("/:id/join", projController.Join)

		taskRoutes(g, projectService)
	}
}

func taskRoutes(e *echo.Group, projectService *services.ProjectService) {
	taskController := controllers.NewTaskController(*projectService)
	g := e.Group("/:proj_id/tasks")
	{
		g.GET("", taskController.Index)
		g.POST("", taskController.Create)
		g.POST("/:id/add-todo", taskController.AddTodo)
	}
}
