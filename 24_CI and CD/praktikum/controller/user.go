package controller

import (
	"ci-cd-go/dto"
	"ci-cd-go/helper"
	"ci-cd-go/model"
	"ci-cd-go/repository"

	"github.com/labstack/echo/v4"
)

type UserController interface{}

type userController struct {
	userRepo repository.UserRepository
}

func NewUserController(uRepo repository.UserRepository) *userController {
	return &userController{
		uRepo,
	}
}

func (u *userController) GetAllUsers(c echo.Context) error {
	users, err := u.userRepo.Find()
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(200, echo.Map{
		"data": users,
	})
}

func (u *userController) CreateUser(c echo.Context) error {
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}
	err = u.userRepo.Create(user)
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	token, err := helper.CreateToken(user.Email)
	if err != nil {
		return c.JSON(401, echo.Map{
			"error": err.Error(),
		})
	}
	uRes := dto.DTOUser{
		user.Email,
		user.Password,
		token,
	}
	return c.JSON(200, echo.Map{
		"data": uRes,
	})

}

//func GetAllUsers(db *gorm.DB) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		users := make([]model.User, 0)
//		err := db.Find(&users).Error
//		if err != nil {
//			return c.JSON(500, echo.Map{
//				"error": err.Error(),
//			})
//		}
//		return c.JSON(200, echo.Map{
//			"data": users,
//		})
//	}
//}
//
//func CreateUser(db *gorm.DB) echo.HandlerFunc {
//	user := model.User{}
//	return func(c echo.Context) error {
//		if err := c.Bind(&user); err != nil {
//			return c.JSON(400, echo.Map{
//				"error": err.Error(),
//			})
//		}
//		err := db.Create(&user).Error
//		if err != nil {
//			return c.JSON(500, echo.Map{
//				"error": err.Error(),
//			})
//		}
//		return c.JSON(200, echo.Map{
//			"data": user,
//		})
//	}
//}
