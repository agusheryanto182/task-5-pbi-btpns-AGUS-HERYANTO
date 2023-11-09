package services

import (
	"errors"

	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/app"
	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/models"
	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/repositories"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	userRepository repositories.UserRepository
	validate       *validator.Validate
}

func (s *UserServiceImpl) RegisterUser(input app.RegisterUserInput) (models.User, error) {
	err := s.validate.Struct(input)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{}
	user.Username = input.Username
	user.Email = input.Email

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)

	newUser, err := s.userRepository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *UserServiceImpl) LoginUser(input app.LoginInput) (models.User, error) {
	err := s.validate.Struct(input)
	if err != nil {
		return models.User{}, err
	}

	email := input.Email
	password := input.Password

	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *UserServiceImpl) IsEmailAvailable(input string) (bool, error) {
	err := s.validate.Struct(input)
	if err != nil {
		return false, err
	}

	email := input
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return false, err
	}
	if user.ID == 0 {
		return true, nil
	}
	return false, nil
}

func (s *UserServiceImpl) IsUsernameAvailable(input string) (bool, error) {
	err := s.validate.Struct(input)
	if err != nil {
		return false, err
	}

	username := input
	user, err := s.userRepository.FindByUsername(username)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}
	return false, nil
}

func (s *UserServiceImpl) GetUserByID(ID int) (models.User, error) {
	user, err := s.userRepository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User not found on that ID")
	}
	return user, nil
}

func (s *UserServiceImpl) UpdateUser(inputID app.GetUserDetailInput, inputData app.FormUpdateUserInput) (models.User, error) {
	err := s.validate.Struct(inputData)
	if err != nil {
		return models.User{}, err
	}

	user, err := s.userRepository.FindByID(inputID.ID)
	if err != nil {
		return user, err
	}

	if user.ID != inputData.User.ID {
		return user, err
	}

	if user.Username == inputData.Username {
		user.Username = inputData.Username
	}

	var isUsernameAvailable string
	isUsernameAvailable = inputData.Username

	checkUsername, _ := s.IsUsernameAvailable(isUsernameAvailable)
	if checkUsername {
		user.Username = inputData.Username
	}

	if user.Email == inputData.Email {
		user.Email = inputData.Email
	}

	var isEmailAvailable string
	isEmailAvailable = inputData.Email

	checkEmail, _ := s.IsEmailAvailable(isEmailAvailable)
	if checkEmail {
		user.Email = inputData.Email
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(inputData.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)

	updatedUser, err := s.userRepository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil

}

func (s *UserServiceImpl) DeleteUser(ID int) error {
	err := s.userRepository.Delete(ID)
	if err != nil {
		return err
	}
	return nil
}

func NewUserService(userRepository repositories.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{userRepository: userRepository, validate: validate}
}
