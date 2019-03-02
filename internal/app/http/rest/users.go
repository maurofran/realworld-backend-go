package rest

import "github.com/labstack/echo"

func installUsersRoutes(g *echo.Group) {
	g.POST("/users/login", loginUser)
	g.POST("/users", registerUser)
	g.GET("/user", currentUser)
	g.PUT("/user", updateUser)
}

func loginUser(ctx echo.Context) error {
	panic("not implemented")
}

func registerUser(ctx echo.Context) error {
	panic("not implemented")
}

func currentUser(ctx echo.Context) error {
	panic("not implemented")
}

func updateUser(ctx echo.Context) error {
	panic("not implemented")
}