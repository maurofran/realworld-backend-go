package rest

import "github.com/labstack/echo"

func installArticlesRoutes(g *echo.Group) {
	g.GET("/articles", getArticles)
	g.GET("/articles/feed", articlesFeed)
	g.GET("/articles/:slug", getArticle)
	g.POST("/articles", createArticle)
	g.PUT("/articles/:slug", updateArticle)
	g.DELETE("/articles/:slug", deleteArticle)
	g.POST("/articles/:slug/comments", addArticleComment)
	g.GET("/articles/:slug/comments", getArticleComments)
	g.DELETE("/articles/:slug/comments/:id", deleteArticleComment)
	g.POST("/articles/:slug/favorite", addArticleToFavorites)
	g.DELETE("/articles/:slug/favorite", removeArticleFromFavorites)
	g.GET("/tags", getTags)
}

func getArticles(ctx echo.Context) error {
	panic("not implemented")
}

func articlesFeed(ctx echo.Context) error {
	panic("not implemented")
}

func getArticle(ctx echo.Context) error {
	panic("not implemented")
}

func createArticle(ctx echo.Context) error {
	panic("not implemented")
}

func updateArticle(ctx echo.Context) error {
	panic("not implemented")
}

func deleteArticle(ctx echo.Context) error {
	panic("not implemented")
}

func addArticleComment(ctx echo.Context) error {
	panic("not implemented")
}

func getArticleComments(ctx echo.Context) error {
	panic("not implemented")
}

func deleteArticleComment(ctx echo.Context) error {
	panic("not implemented")
}

func addArticleToFavorites(ctx echo.Context) error {
	panic("not implemented")
}

func removeArticleFromFavorites(ctx echo.Context) error {
	panic("not implemented")
}

func getTags(ctx echo.Context) error {
	panic("not implemented")
}