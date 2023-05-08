package services

import (
	"user_list_go/app/models"
	"user_list_go/app/repositories"
)

type UserService struct {
	userRepo repositories.UserRepositoryInterface
}

func NewUserService(uRepo repositories.UserRepositoryInterface) *UserService {
	return &UserService{
		userRepo: uRepo,
	}
}

func (uRepo *UserService) GetUsers() ([]models.User, error) {
	return uRepo.userRepo.GetUsers()
}

func (uRepo *UserService) GetUser(id int) (models.User, error) {
	return uRepo.userRepo.GetUser(id)
}

func (uRepo *UserService) CreateUser(user models.User) error {
	return uRepo.userRepo.CreateUser(user)
}

func (uRepo *UserService) UpdateUser(id int, user models.User) error {
	return uRepo.userRepo.UpdateUser(id, user)
}

func (uRepo *UserService) DeleteUser(id int) error {
	return uRepo.userRepo.DeleteUser(id)
}
