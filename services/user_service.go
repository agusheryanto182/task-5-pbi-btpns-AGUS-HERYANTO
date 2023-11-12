package services

import (
	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/app"
	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/models"
)

type UserService interface {
	RegisterUser(input app.RegisterUserInput) (models.User, error)
	LoginUser(input app.LoginInput) (models.User, error)
	IsEmailAvailable(input string) (bool, error)
	IsUsernameAvailable(input string) (bool, error)
	GetUserByID(ID int) (models.User, error)
	UpdateUser(userID int, inputData app.FormUpdateUserInput) (models.User, error)
	DeleteUser(ID int) error
}
