package routes

import (
	"belajar-go-echo/controller"
	"belajar-go-echo/repository"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewRoute(app *echo.Echo, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userController := controller.NewUserController(userRepository)

	app.GET("/users", userController.GetAllUsers)
	app.POST("/users", userController.CreateUser)
}
