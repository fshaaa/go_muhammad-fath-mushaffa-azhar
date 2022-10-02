package database

import (
	"problem-2/config"
	"problem-2/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User

	err := config.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
