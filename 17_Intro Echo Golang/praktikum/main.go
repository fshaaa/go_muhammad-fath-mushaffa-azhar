package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

func GetUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	idUser, _ := strconv.Atoi(c.Param("id"))
	for _, u := range users {
		if u.Id == idUser {
			user = u
			break
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by user id",
		"user":    user,
	})
}

func DeleteUserController(c echo.Context) error {
	idUser, _ := strconv.Atoi(c.Param("id"))
	for i, u := range users {
		if u.Id == idUser {
			if len(users) == 0 {
				users = []User{}
			} else {
				users[i] = users[0]
				users = users[1:]
			}
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user by user id",
		"user":    users[0],
	})
}

func UpdateUserController(c echo.Context) error {
	nameUser := c.QueryParam("name")
	emailUser := c.QueryParam("email")
	passUser := c.QueryParam("password")
	var key int
	idUser, _ := strconv.Atoi(c.Param("id"))
	for i, u := range users {
		if idUser == u.Id {
			if nameUser != "" {
				users[i].Name = nameUser
			}
			if emailUser != "" {
				users[i].Email = emailUser
			}
			if passUser != "" {
				users[i].Password = passUser
			}
			key = i
			break
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user by user id",
		"user":    users[key],
	})
}

func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	user.Name = c.QueryParam("name")
	user.Email = c.QueryParam("email")
	user.Password = c.QueryParam("password")
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"user":    user,
	})
}

func main() {
	e := echo.New()

	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.GET("/users/:id", GetUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	e.Logger.Fatal(e.Start(":8000"))
}
