package database

import (
	"docker-problem/config"
	"docker-problem/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User

	err := config.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
