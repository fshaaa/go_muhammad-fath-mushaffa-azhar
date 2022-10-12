package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"praktikum/config"
	"praktikum/lib/database"
	"praktikum/models"
)

func GetBooksController(c echo.Context) error {
	books, err := database.GetBooks()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

func GetBookController(c echo.Context) error {
	id := c.Param("id")
	var book models.Book

	err := config.DB.Find(&book, id).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book by id",
		"book":    book,
	})
}

func CreateBookController(c echo.Context) error {
	var book models.Book
	c.Bind(&book)

	err := config.DB.Create(&book).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create book",
		"book":    book,
	})
}

func UpdateBookController(c echo.Context) error {
	id := c.Param("id")
	var book models.Book
	c.Bind(&book)

	err := config.DB.Where("id = ?", id).Updates(book).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book",
		"book":    book,
	})
}

func DeleteBookController(c echo.Context) error {
	id := c.Param("id")
	var books []models.Book

	err := config.DB.Delete(&books, id).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book",
		"book":    books,
	})
}
