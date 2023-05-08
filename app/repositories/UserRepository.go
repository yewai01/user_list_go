package repositories

import "user_list_go/app/models"

type UserRepository struct {
	Users []models.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (uRepo *UserRepository) GetUsers() ([]models.User, error) {
	return []models.User{
		{ID: 1, Name: "Ye Wai", Birthday: 1994, Gender: "Female"},
		{ID: 2, Name: "The Godfather", Birthday: 1972, Gender: "Male"},
		{ID: 3, Name: "The Dark Knight", Birthday: 2008, Gender: "Male"},
	}, nil
}

func (uRepo *UserRepository) GetUser(id int) (models.User, error) {
	return models.User{}, nil
}

func (uRepo *UserRepository) CreateUser(user models.User) error {
	return nil
}

func (uRepo *UserRepository) UpdateUser(id int, user models.User) error {
	return nil
}

func (uRepo *UserRepository) DeleteUser(id int) error {
	return nil
}
