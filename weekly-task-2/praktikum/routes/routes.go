package routes

import (
	"github.com/labstack/echo/v4"
	"praktikum-weekly-task-2/controllers"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/items", controllers.GetItemsController)
	e.GET("/items/:id", controllers.GetItemController)
	e.POST("/items", controllers.CreateItemController)
	e.PUT("/items/:id", controllers.UpdateItemController)
	e.DELETE("/items/:id", controllers.DeleteItemController)
	e.GET("/items/category/:category_id", controllers.GetCategoryController)
	//e.GET("/items?keyword=item_name", controllers.GetItemNameController)

	return e
}
