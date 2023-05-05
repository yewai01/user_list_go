package main

import (
	"github.com/labstack/echo/v4"
	// "github.com/yewai01/user_list_go/handlers"
	// "github.com/yewai01/user_list_go/repositories"
	// "github.com/yewai01/user_list_go/services"
)

func main() {

	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	e := echo.New()

	e.GET("/users", userHandler.GetUsers)
	e.GET("/users/:id", userHandler.GetUser)
	e.POST("/users", userHandler.CreateUser)
	e.PATCH("/users/:id", userHandler.UpdateUser)
	e.DELETE("/users", userHandler.DeleteAllUsers)
	e.DELETE("/users/:id", userHandler.DeleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}
