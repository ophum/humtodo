package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ophum/humtodo/pkg/services"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// +gen-ts-entity
type SignUpRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// +gen-ts-entity
type SignUpResponse struct {
	Token string `json:"token"`
}

func (c *AuthController) SignUp(ctx echo.Context) error {
	req := SignUpRequest{}
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	token, err := c.authService.SignUp(req.Name, req.Password)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	return ctx.JSON(http.StatusCreated, SignUpResponse{
		Token: token,
	})
}

// +gen-ts-entity
type SignInRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// +gen-ts-entity
type SignInResponse struct {
	Token string `json:"token"`
}

func (c *AuthController) SignIn(ctx echo.Context) error {
	req := SignInRequest{}
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	token, err := c.authService.SignIn(req.Name, req.Password)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, SignInResponse{
		Token: token,
	})
}

func (c *AuthController) SignOut(ctx echo.Context) error {
	return nil
}

// +gen-ts-entity
type VerifyRequest struct {
	Token string `json:"token"`
}

// +gen-ts-entity
type VerifyResponse struct {
	Token string `json:"token"`
}

func (c *AuthController) Verify(ctx echo.Context) error {
	req := VerifyRequest{}
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	token, err := c.authService.Verify(req.Token)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, VerifyResponse{
		Token: token,
	})
}
