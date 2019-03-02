package rest

import "github.com/labstack/echo"

func installProfileRoutes(g *echo.Group) {
	g.GET("/profiles/:username", getProfile)
	g.POST("/profiles/:username/follow", followProfile)
	g.DELETE("/profiles/:username/follow", unfollowProfile)
}

func getProfile(ctx echo.Context) error {
	panic("not implemented")
}

func followProfile(ctx echo.Context) error {
	panic("not implemented")
}

func unfollowProfile(ctx echo.Context) error {
	panic("not implemented")
}