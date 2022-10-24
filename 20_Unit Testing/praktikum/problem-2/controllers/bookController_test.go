package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"praktikum-unit-testing/models"
	"regexp"
	"testing"
)

func TestGetBooksController(t *testing.T) {
	mock := Mocking(t)

	var testCases = []struct {
		Name           string
		Path           string
		ExpectCode     int
		HasReturnBody  bool
		ExpectResponse map[string]interface{}
		DatabaseRows   *sqlmock.Rows
	}{
		{
			Name:          "get all books",
			Path:          "/books",
			ExpectCode:    http.StatusOK,
			HasReturnBody: true,
			ExpectResponse: map[string]interface{}{
				"books": []interface{}{
					map[string]interface{}{
						"CreatedAt":     "0001-01-01T00:00:00Z",
						"UpdatedAt":     "0001-01-01T00:00:00Z",
						"DeletedAt":     nil,
						"ID":            float64(1),
						"title":         "astronomi",
						"author":        "azhar",
						"publisherYear": float64(1990),
						"totalPage":     float64(100),
					},
				},
				"message": "success get all books",
			},
			DatabaseRows: sqlmock.NewRows([]string{"ID", "PublisherYear", "TotalPage", "Title", "Author"}).
				AddRow(1, 1990, 100, "astronomi", "azhar"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `books` WHERE `books`.`deleted_at` IS NULL")).
				WillReturnRows(testCase.DatabaseRows)

			r := httptest.NewRequest(http.MethodGet, "/", nil)
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)

			err := GetBooksController(ctx)
			assert.NoError(t, err)
			assert.Equal(t, testCase.ExpectCode, w.Result().StatusCode)
			//assert.Equal(t, testCase.ExceptBody, w.Result().Body)

			if testCase.HasReturnBody {
				var response map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&response)
				assert.NoError(t, err)
				assert.Equal(t, testCase.ExpectResponse, response)
			}
		})
	}
}

func TestGetBookController(t *testing.T) {
	mock := Mocking(t)

	var testCases = []struct {
		Name           string
		Path           string
		ExpectCode     int
		HasReturnBody  bool
		ExpectResponse map[string]interface{}
		DatabaseRows   *sqlmock.Rows
	}{
		{
			Name:          "get book by id",
			Path:          "/books/:id",
			ExpectCode:    200,
			HasReturnBody: true,
			ExpectResponse: map[string]interface{}{
				"book": map[string]interface{}{
					"CreatedAt":     "0001-01-01T00:00:00Z",
					"UpdatedAt":     "0001-01-01T00:00:00Z",
					"DeletedAt":     nil,
					"ID":            float64(1),
					"title":         "astronomi",
					"author":        "azhar",
					"publisherYear": float64(1990),
					"totalPage":     float64(100),
				},
				"message": "success get book by id",
			},
			DatabaseRows: sqlmock.NewRows([]string{"ID", "PublisherYear", "TotalPage", "Title", "Author"}).
				AddRow(1, 1990, 100, "astronomi", "azhar"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `books` WHERE `books`.`id` = ? AND `books`.`deleted_at` IS NULL")).
				WithArgs("1").
				WillReturnRows(testCase.DatabaseRows)

			r := httptest.NewRequest(http.MethodGet, "/", nil)
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")

			err := GetBookController(ctx)
			assert.NoError(t, err)
			assert.Equal(t, testCase.ExpectCode, w.Result().StatusCode)

			if testCase.HasReturnBody {
				var response map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&response)
				assert.NoError(t, err)
				assert.Equal(t, testCase.ExpectResponse, response)
			}
		})
	}
}

func TestCreateBookController(t *testing.T) {
	mock := Mocking(t)

	var testCases = []struct {
		Name           string
		Path           string
		ExpectCode     int
		HasReturnBody  bool
		ExpectResponse map[string]interface{}
		DatabaseRows   *sqlmock.Rows
		BodyArgs       map[string]interface{}
	}{
		{
			Name:          "create book",
			Path:          "/books",
			ExpectCode:    200,
			HasReturnBody: true,
			ExpectResponse: map[string]interface{}{
				"message": "success create book",
				"book": map[string]interface{}{
					"CreatedAt":     "0001-01-01T00:00:00Z",
					"UpdatedAt":     "0001-01-01T00:00:00Z",
					"DeletedAt":     nil,
					"ID":            float64(1),
					"title":         "astronomi",
					"author":        "azhar",
					"publisherYear": float64(1990),
					"totalPage":     float64(100),
				},
			},
			DatabaseRows: sqlmock.NewRows([]string{"ID", "PublisherYear", "TotalPage", "Title", "Author"}).
				AddRow(1, 1990, 100, "astronomi", "azhar"),
			BodyArgs: map[string]interface{}{
				"body": models.Book{
					Title:         "astronomi",
					Author:        "azhar",
					PublisherYear: 1990,
					TotalPage:     10,
				},
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta(" INSERT INTO `books` (`title`,`author`,`publisher_year`,`total_page`) VALUES (?,?,?,?)")).
				WithArgs("", "", 0, 0).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			jsonBody, _ := json.Marshal(testCase.BodyArgs["body"])
			r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonBody))
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/")

			err := CreateBookController(ctx)
			assert.NoError(t, err)
			assert.Equal(t, testCase.ExpectCode, w.Result().StatusCode)

			if testCase.HasReturnBody {
				var response map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&response)
				assert.NoError(t, err)
				assert.Equal(t, testCase.ExpectResponse["message"], response["message"])
			}
		})
	}
}

func TestDeleteBookController(t *testing.T) {
	mock := Mocking(t)

	var testCases = []struct {
		Name           string
		Path           string
		ExpectCode     int
		HasReturnBody  bool
		ExpectResponse map[string]interface{}
		DatabaseRows   *sqlmock.Rows
	}{
		{
			Name:          "delete book by id",
			Path:          "/books",
			ExpectCode:    200,
			HasReturnBody: true,
			ExpectResponse: map[string]interface{}{
				"message": "success delete book",
				"book": map[string]interface{}{
					"CreatedAt":     "0001-01-01T00:00:00Z",
					"UpdatedAt":     "0001-01-01T00:00:00Z",
					"DeletedAt":     nil,
					"ID":            float64(0),
					"title":         "",
					"author":        "",
					"publisherYear": float64(0),
					"totalPage":     float64(0),
				},
			},
			DatabaseRows: sqlmock.NewRows([]string{"ID", "PublisherYear", "TotalPage", "Title", "Author"}).
				AddRow(1, 1990, 100, "astronomi", "azhar"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `books` WHERE `books`.`id` = ?")).
				WithArgs("1").
				WillReturnResult(sqlmock.NewResult(1, 0)).
				WillReturnError(nil)
			mock.ExpectCommit()

			r := httptest.NewRequest(http.MethodDelete, "/", nil)
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")

			err := DeleteBookController(ctx)
			assert.NoError(t, err)
			assert.Equal(t, testCase.ExpectCode, w.Result().StatusCode)

			if testCase.HasReturnBody {
				var response map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&response)
				assert.NoError(t, err)
				assert.Equal(t, testCase.ExpectResponse, response)
			}
		})
	}
}

func TestUpdateBookController(t *testing.T) {
	mock := Mocking(t)

	var testCases = []struct {
		Name           string
		Path           string
		ExpectCode     int
		HasReturnBody  bool
		ExpectResponse map[string]interface{}
		DatabaseRows   *sqlmock.Rows
	}{
		{
			Name:          "update book by id",
			Path:          "/books",
			ExpectCode:    200,
			HasReturnBody: true,
			ExpectResponse: map[string]interface{}{
				"message": "success update book",
				"book": map[string]interface{}{
					"CreatedAt":     "0001-01-01T00:00:00Z",
					"UpdatedAt":     "0001-01-01T00:00:00Z",
					"DeletedAt":     nil,
					"ID":            float64(0),
					"title":         "",
					"author":        "",
					"publisherYear": float64(0),
					"totalPage":     float64(0),
				},
			},
			DatabaseRows: sqlmock.NewRows([]string{"ID", "PublisherYear", "TotalPage", "Title", "Author"}).
				AddRow(1, 1990, 100, "astronomi", "azhar"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta("UPDATE `books` SET `updated_at`=? WHERE id = ?")).
				WithArgs(AnyTime{}, "1").
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			r := httptest.NewRequest(http.MethodPut, "/", nil)
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")

			err := UpdateBookController(ctx)
			assert.NoError(t, err)
			assert.Equal(t, testCase.ExpectCode, w.Result().StatusCode)

			if testCase.HasReturnBody {
				var response map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&response)
				assert.NoError(t, err)
				//assert.Equal(t, t, testCase.ExpectResponse["message"], response["message"])
			}
		})
	}
}
