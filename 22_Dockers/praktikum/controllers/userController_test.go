package controllers

import (
	"bytes"
	"database/sql/driver"
	"docker-problem/config"
	"docker-problem/models"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"
)

type AnyTime struct {
}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func Mocking(t *testing.T) sqlmock.Sqlmock {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	dbGorm, _ := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}))

	config.DB = dbGorm

	return mock
}

func TestGetUsersController(t *testing.T) {
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
			Name:          "get all user",
			Path:          "/users",
			ExpectCode:    http.StatusOK,
			HasReturnBody: true,
			ExpectResponse: map[string]interface{}{
				"message": "success get all user",
				"user": []interface{}{
					map[string]interface{}{
						"CreatedAt": "0001-01-01T00:00:00Z",
						"UpdatedAt": "0001-01-01T00:00:00Z",
						"DeletedAt": nil,
						"ID":        float64(1),
						"name":      "azhar",
						"email":     "azhar@gmail.com",
						"password":  "azhargantenk",
					},
				},
			},
			DatabaseRows: sqlmock.NewRows([]string{"ID", "name", "email", "password"}).
				AddRow(1, "azhar", "azhar@gmail.com", "azhargantenk"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL")).
				WillReturnRows(testCase.DatabaseRows)

			r := httptest.NewRequest(http.MethodGet, "/", nil)
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)

			err := GetUsersController(ctx)
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

func TestGetUserController(t *testing.T) {
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
			Name:          "get user by id",
			Path:          "/users/:id",
			ExpectCode:    200,
			HasReturnBody: true,
			ExpectResponse: map[string]interface{}{
				"message": "success get user by id",
				"user": map[string]interface{}{
					"CreatedAt": "0001-01-01T00:00:00Z",
					"UpdatedAt": "0001-01-01T00:00:00Z",
					"DeletedAt": nil,
					"ID":        float64(1),
					"name":      "azhar",
					"email":     "azhar@gmail.com",
					"password":  "azhargantenk",
				},
			},
			DatabaseRows: sqlmock.NewRows([]string{"ID", "name", "email", "password"}).
				AddRow("1", "azhar", "azhar@gmail.com", "azhargantenk"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`id` = ? AND `users`.`deleted_at` IS NULL")).
				WithArgs("1").
				WillReturnRows(testCase.DatabaseRows)

			r := httptest.NewRequest(http.MethodGet, "/", nil)
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")

			err := GetUserController(ctx)
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

func TestCreateUserController(t *testing.T) {
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
			Name:          "create user",
			Path:          "/users",
			ExpectCode:    200,
			HasReturnBody: true,
			ExpectResponse: map[string]interface{}{
				"message": "success create user",
				"user": map[string]interface{}{
					"CreatedAt": ".",
					"UpdatedAt": ".",
					"DeletedAt": nil,
					"ID":        float64(1),
					"name":      "azhar",
					"email":     "azhar@gmail.com",
					"password":  "azhargantenk",
				},
			},
			DatabaseRows: sqlmock.NewRows([]string{"name", "email", "password"}).
				AddRow("azhar", "azhar@gmail.com", "azhargantenk"),
			BodyArgs: map[string]interface{}{
				"body": models.User{
					Name:     "azhar",
					Email:    "azhar@gmail.com",
					Password: "azhargantenk",
				},
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`name`,`email`,`password`) VALUES (?,?,?)")).
				WithArgs("", "", "").
				WillReturnResult(sqlmock.NewResult(0, 1))
			mock.ExpectCommit()

			jsonBody, _ := json.Marshal(testCase.BodyArgs["body"])
			r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonBody))
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/")

			err := CreateUserController(ctx)
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

func TestDeleteUserController(t *testing.T) {
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
			Name:          "delete user by id",
			Path:          "/users/:id",
			ExpectCode:    200,
			HasReturnBody: true,
			ExpectResponse: map[string]interface{}{
				"message": "success delete user by id",
				"user": map[string]interface{}{
					"CreatedAt": "0001-01-01T00:00:00Z",
					"UpdatedAt": "0001-01-01T00:00:00Z",
					"DeletedAt": nil,
					"ID":        float64(0),
					"name":      "",
					"email":     "",
					"password":  "",
				},
			},
			DatabaseRows: sqlmock.NewRows([]string{"ID", "name", "email", "password"}).
				AddRow("1", "azhar", "azhar@gmail.com", "azhargantenk"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `users` WHERE `users`.`id` = ?")).
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

			err := DeleteUserController(ctx)
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

func TestUpdateUserController(t *testing.T) {
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
			Name:          "update user by id",
			Path:          "/users/:id",
			ExpectCode:    200,
			HasReturnBody: true,
			ExpectResponse: map[string]interface{}{
				"message": "success update user by id",
				"user": map[string]interface{}{
					"CreatedAt": "0001-01-01T00:00:00Z",
					"UpdatedAt": ".",
					"DeletedAt": nil,
					"ID":        float64(0),
					"name":      "",
					"email":     "",
					"password":  "",
				},
			},
			DatabaseRows: sqlmock.NewRows([]string{"ID", "name", "email", "password"}).
				AddRow("1", "azhar", "azhar@gmail.com", "azhargantenk"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta("UPDATE `users` SET `updated_at`=? WHERE id = ?")).
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

			err := UpdateUserController(ctx)
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
