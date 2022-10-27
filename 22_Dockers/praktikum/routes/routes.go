package routes

import (
	"docker-problem/constants"
	"docker-problem/controllers"
	"docker-problem/middleware"
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	middleware.LogMiddleware(e)

	userJwt := e.Group("/users")
	userJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	userJwt.GET("", controllers.GetUsersController)
	userJwt.GET("/:id", controllers.GetUserController)
	userJwt.PUT("/:id", controllers.UpdateUserController)
	userJwt.DELETE("/:id", controllers.DeleteUserController)
	e.POST("/users", controllers.CreateUserController)

	bookJwt := e.Group("/books")
	bookJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	bookJwt.POST("", controllers.CreateBookController)
	bookJwt.PUT("/:id", controllers.UpdateBookController)
	bookJwt.DELETE("/:id", controllers.DeleteBookController)
	e.GET("/books", controllers.GetBooksController)
	e.GET("/books/:id", controllers.GetBookController)

	return e
}
