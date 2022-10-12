package database

import (
	"praktikum/config"
	"praktikum/models"
)

func GetBooks() (interface{}, error) {
	var book []models.Book

	err := config.DB.Find(&book).Error
	if err != nil {
		return nil, err
	}
	return book, nil
}
