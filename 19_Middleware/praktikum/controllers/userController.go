package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"praktikum/config"
	"praktikum/lib/database"
	"praktikum/middleware"
	"praktikum/models"
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

func LoginUserController(c echo.Context) error {
	var user models.User
	c.Bind(&user)

	err := config.DB.Where("email = ? && password = ?", user.Email, user.Password).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	token, err := middleware.CreateToken(user.ID, user.Name)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	userRespon := models.UserResponse{
		user.Model,
		user.Name,
		user.Email,
		token,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"user":    userRespon,
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
	var user models.User

	err := config.DB.Delete(&user, id).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user by id",
		"user":    user,
	})
}
