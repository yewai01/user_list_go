package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"user_list_go/app/models"
	"user_list_go/app/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type userHandler struct {
	service services.UserServiceInterface
}

type errorResponse struct {
	Error string `json:"error"`
}

type successResponse struct {
	Message string `json:"message"`
}

func NewUserHandler(us services.UserServiceInterface) *userHandler {
	return &userHandler{service: us}
}

func (handler *userHandler) GetUsers(c echo.Context) error {
	users, err := handler.service.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse{
			Error: err.Error(),
		})
	}
	if len(users) == 0 {
		return c.JSON(http.StatusNotFound, errorResponse{
			Error: "there are no records",
		})
	}

	return c.JSON(http.StatusOK, users)
}

func (handler *userHandler) GetUser(c echo.Context) error {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	user, err := handler.service.GetUser(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, errorResponse{
				Error: err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, errorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, user)
}

func (handler *userHandler) CreateUser(c echo.Context) error {
	user := models.User{}
	defer c.Request().Body.Close()

	b, _ := ioutil.ReadAll(c.Request().Body)
	json.Unmarshal(b, &user)

	err := handler.service.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse{
			Error: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, successResponse{
		Message: "create success",
	})
}

func (handler *userHandler) UpdateUser(c echo.Context) error {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	user := models.User{}
	defer c.Request().Body.Close()

	b, _ := ioutil.ReadAll(c.Request().Body)
	json.Unmarshal(b, &user)

	err := handler.service.UpdateUser(id, user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, errorResponse{
				Error: err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, errorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, successResponse{
		Message: "update success",
	})
}

func (handler *userHandler) DeleteUser(c echo.Context) error {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	err := handler.service.DeleteUser(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, errorResponse{
				Error: err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, errorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, successResponse{
		Message: "delete success",
	})
}
