package routes

import (
	"user_list_go/app/handlers"
	"user_list_go/app/repositories"
	"user_list_go/app/services"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	e := echo.New()

	e.GET("/users", userHandler.GetUsers)
	e.GET("/users/:id", userHandler.GetUser)
	e.POST("/users", userHandler.CreateUser)
	e.PATCH("/users/:id", userHandler.UpdateUser)
	e.DELETE("/users/:id", userHandler.DeleteUser)

	return e
}
