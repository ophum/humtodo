package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

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
