package database

import (
	"praktikum-unit-testing/config"
	"praktikum-unit-testing/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User

	err := config.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
