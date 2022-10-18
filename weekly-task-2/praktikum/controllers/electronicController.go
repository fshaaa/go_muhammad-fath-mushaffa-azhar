package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"praktikum-weekly-task-2/database"
	"praktikum-weekly-task-2/helper"
	"praktikum-weekly-task-2/models"
	"strconv"
)

func GetItemsController(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	if keyword != "" {
		return GetItemNameController(c)
	}

	item, err := database.GetItems()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all item",
		"item":    item,
	})
}

func GetItemController(c echo.Context) error {
	id := c.Param("id")
	item, err := database.GetItem(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get a item",
		"item":    item,
	})
}

func CreateItemController(c echo.Context) error {
	var item models.Electronics
	c.Bind(&item)
	item.ID = helper.CreateUUID()

	err := database.CreateItem(&item)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create item",
		"item":    item,
	})
}

func UpdateItemController(c echo.Context) error {
	id := c.Param("id")
	var item models.Electronics
	c.Bind(&item)

	err := database.UpdateItem(id, &item)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update item",
		"item":    item,
	})
}

func DeleteItemController(c echo.Context) error {
	id := c.Param("id")

	err := database.DeleteItem(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete item",
	})
}

func GetCategoryController(c echo.Context) error {
	cat_id, _ := strconv.Atoi(c.Param("category_id"))

	items, err := database.GetCategory(cat_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get item by category id",
		"item":    items,
	})
}

func GetItemNameController(c echo.Context) error {
	keyword := c.QueryParam("keyword")

	items, err := database.GetItemName(keyword)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get item by name",
		"item":    items,
	})
}
