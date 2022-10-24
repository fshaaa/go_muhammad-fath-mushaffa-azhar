package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"praktikum-unit-testing/config"
	"praktikum-unit-testing/lib/database"
	"praktikum-unit-testing/models"
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

	err := config.DB.Omit("created_at", "updated_at", "deleted_at").Save(&book).Error
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
	var book models.Book

	err := config.DB.Unscoped().Delete(&book, id).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book",
		"book":    book,
	})
}
