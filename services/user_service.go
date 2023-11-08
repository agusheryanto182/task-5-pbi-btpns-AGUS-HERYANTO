package services

import "github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/models"

type UserService interface {
	RegisterUser(input models.RegisterUserInput) (models.User, error)
	Login(input models.LoginInput) (models.User, error)
	IsEmailAvailable(input string) (bool, error)
	IsUsernameAvailable(input string) (bool, error)
	GetUserByID (ID int) (models.User, error)
	UpdateUser(inputID models.GetUserDetailInput, inputData models.FormUpdateUserInput) (models.User, error)
	DeleteUser(ID int) error
}
