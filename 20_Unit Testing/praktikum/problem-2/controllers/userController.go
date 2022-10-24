package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"praktikum-unit-testing/config"
	"praktikum-unit-testing/middleware"
	"praktikum-unit-testing/models"
)

func GetUsersController(c echo.Context) error {
	var users []models.User

	err := config.DB.Find(&users).Error

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
	var user models.User

	err := config.DB.Find(&user, id).Error
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
	
	err := config.DB.Omit("created_at", "updated_at", "deleted_at").Save(&user).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	_, err = middleware.CreateToken(user.ID, user.Name)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
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

	err := config.DB.Where("id = ?", id).Updates(&user).Error
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
	var user models.User

	err := config.DB.Unscoped().Delete(&user, id).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user by id",
		"user":    user,
	})
}
