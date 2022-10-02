package database

import (
	"problem-2/config"
	"problem-2/models"
)

func GetBooks() (interface{}, error) {
	var book []models.Books

	err := config.DB.Find(&book).Error
	if err != nil {
		return nil, err
	}
	return book, nil
}
