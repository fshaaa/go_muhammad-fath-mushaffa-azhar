package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"problem-2/config"
	"problem-2/lib/database"
	"problem-2/models"
)

func GetUsersController(c echo.Context) error {
	users, err := database.GetUsers()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all user",
		"user":    users,
	})
}

func GetUserController(c echo.Context) error {
	id := c.Param("id")
	var user []models.User

	err := config.DB.Find(&user).Where("id = ?", id).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by id",
		"user":    user,
	})
}

func CreateUserController(c echo.Context) error {
	var user models.User
	c.Bind(&user)

	err := config.DB.Create(&user).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"user":    user,
	})
}

func UpdateUserController(c echo.Context) error {
	id := c.Param("id")
	var user models.User
	c.Bind(&user)

	err := config.DB.Where("id = ?", id).Updates(user).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user by id",
		"user":    user,
	})
}

func DeleteUserController(c echo.Context) error {
	id := c.Param("id")
	var user []models.User

	err := config.DB.Delete(user).Where("id = ?", id).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user by id",
		"user":    user,
	})
}
