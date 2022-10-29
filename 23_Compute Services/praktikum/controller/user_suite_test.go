package controller

import (
	"fshaaa-rest-api/model"
	"fshaaa-rest-api/repository"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type suiteUsers struct {
	suite.Suite
	control *userController
	mock    *repository.MockUserRepository
}

func (s *suiteUsers) SetupSuite() {
	mock := &repository.MockUserRepository{}
	s.mock = mock

	s.control = &userController{
		userRepo: s.mock,
	}
}

func (s suiteUsers) TestGetAllUsers() {
	s.mock.On("Find").Return([]model.User{
		{
			Email:    "azhar@gmail.com",
			Password: "azhargantenk",
		},
	}, nil)

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		HasReturnBody      bool
		ExpectedBody       model.User
	}{
		{
			"success",
			http.StatusOK,
			"GET",
			true,
			model.User{
				Email:    "azhar@gmail.com",
				Password: "azhargantenk",
			},
		},
	}
	for _, testcase := range testCases {
		s.T().Run(testcase.Name, func(t *testing.T) {
			r := httptest.NewRequest(testcase.Method, "/", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)

			err := s.control.GetAllUsers(ctx)
			s.NoError(err)

			s.Equal(testcase.ExpectedStatusCode, w.Result().StatusCode)
		})
	}
}

func (s suiteUsers) TestCreateUser() {
	mockUser := model.User{
		Email:    "",
		Password: "",
	}
	s.mock.On("Create", mockUser).Return(nil)

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		HasReturnBody      bool
		ExpectedBody       model.User
	}{
		{
			"success",
			http.StatusOK,
			"POST",
			true,
			model.User{
				Email:    "azhar@gmail.com",
				Password: "azhargantenk",
			},
		},
	}
	for _, testcase := range testCases {
		s.T().Run(testcase.Name, func(t *testing.T) {
			r := httptest.NewRequest(testcase.Method, "/", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)

			err := s.control.CreateUser(ctx)
			s.NoError(err)

			s.Equal(testcase.ExpectedStatusCode, w.Result().StatusCode)
		})
	}
}

func TestSuiteUser(t *testing.T) {
	suite.Run(t, new(suiteUsers))
}
