package repositories

import (
	"user_list_go/app/models"
	"user_list_go/bootstrap"
)

type UserRepository struct {
	Users []models.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (uRepo *UserRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	res := bootstrap.DB.Find(&users)
	if res.Error != nil {
		return users, res.Error
	}
	return users, nil
}

func (uRepo *UserRepository) GetUser(id int) (models.User, error) {
	var user models.User

	res := bootstrap.DB.First(&user, id)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func (uRepo *UserRepository) CreateUser(user models.User) error {
	res := bootstrap.DB.Create(&user)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (uRepo *UserRepository) UpdateUser(id int, user models.User) error {
	var oldUser models.User
	res := bootstrap.DB.First(&oldUser, id)
	if res.Error != nil {
		return res.Error
	}
	res = res.Updates(&user)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (uRepo *UserRepository) DeleteUser(id int) error {
	var user models.User
	res := bootstrap.DB.First(&user, id)
	if res.Error != nil {
		return res.Error
	}
	res = res.Delete(&user, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
