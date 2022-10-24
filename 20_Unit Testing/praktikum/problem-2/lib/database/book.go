package database

import (
	"praktikum-unit-testing/config"
	"praktikum-unit-testing/models"
)

func GetBooks() (interface{}, error) {
	var book []models.Book

	err := config.DB.Find(&book).Error
	if err != nil {
		return nil, err
	}
	return book, nil
}
