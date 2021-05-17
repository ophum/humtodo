package routes

import (
	"github.com/labstack/echo"
	"github.com/ophum/humtodo/pkg/controllers"
	"github.com/ophum/humtodo/pkg/repositories"
	"github.com/ophum/humtodo/pkg/services"
)

func Init() *echo.Echo {
	e := echo.New()

	userRepo := repositories.NewUserRepositoryInMemory()

	authRoutes(e, userRepo)

	return e
}

func authRoutes(e *echo.Echo, userRepo repositories.UserRepository) {
	authService := services.NewAuthService([]byte("test"), &userRepo)
	authController := controllers.NewAuthController(*authService)
	g := e.Group("auth")
	{
		g.POST("/sign-in", authController.SignIn)
		g.POST("/sign-up", authController.SignUp)
		g.POST("/verify", authController.Verify)
	}
}
