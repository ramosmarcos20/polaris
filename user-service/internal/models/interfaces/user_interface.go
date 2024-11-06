package interfaces

import "polaris/user-service/internal/models/entities"

type UserInterface interface {
	GetALl() ([]entities.User, error)
	CreateUser(user *entities.User) error
	GetUserById(id uint) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	GetUserByUsername(username string) (*entities.User, error)
	UpdateUser(user *entities.User) error
	Delete(id uint) error
}
