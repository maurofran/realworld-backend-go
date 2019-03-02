package rest

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Api() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Skipper:    skipIfLogin,
		SigningKey: []byte("secret"), // TODO signing key from environment variable
		AuthScheme: "Token",
	}))

	group := e.Group("api")

	installUsersRoutes(group)
	installProfileRoutes(group)
	installArticlesRoutes(group)

	return e
}

func skipIfLogin(ctx echo.Context) bool {
	// Ignore login for authentication.
	return "/api/users/login" == ctx.Request().RequestURI
}
