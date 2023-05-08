package handlers

import (
	"net/http"
	"user_list_go/app/services"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	service services.UserServiceInterface
}

func NewUserHandler(us services.UserServiceInterface) *userHandler {
	return &userHandler{service: us}
}

func (handler *userHandler) GetUsers(c echo.Context) error {
	users, err := handler.service.GetUsers()
	if err != nil {
	}

	return c.JSON(http.StatusOK, users)
}

func (handler *userHandler) GetUser(c echo.Context) error {
	users, err := handler.service.GetUsers()
	if err != nil {
	}

	return c.JSON(http.StatusOK, users)
}

func (handler *userHandler) CreateUser(c echo.Context) error {
	users, err := handler.service.GetUsers()
	if err != nil {
	}

	return c.JSON(http.StatusOK, users)
}

func (handler *userHandler) DeleteUser(c echo.Context) error {
	users, err := handler.service.GetUsers()
	if err != nil {
	}

	return c.JSON(http.StatusOK, users)
}

func (handler *userHandler) DeleteAllUsers(c echo.Context) error {
	users, err := handler.service.GetUsers()
	if err != nil {
	}

	return c.JSON(http.StatusOK, users)
}

func (handler *userHandler) UpdateUser(c echo.Context) error {
	users, err := handler.service.GetUsers()
	if err != nil {
	}

	return c.JSON(http.StatusOK, users)
}
