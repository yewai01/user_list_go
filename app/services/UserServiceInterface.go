package services

import "user_list_go/app/models"

type UserServiceInterface interface {
	GetUsers() ([]models.User, error)
	GetUser(id int) (models.User, error)
	CreateUser(user models.User) error
	UpdateUser(id int, user models.User) error
	DeleteUser(id int) error
}
