package database

import (
	"docker-problem/config"
	"docker-problem/models"
)

func GetBooks() (interface{}, error) {
	var book []models.Book

	err := config.DB.Find(&book).Error
	if err != nil {
		return nil, err
	}
	return book, nil
}
