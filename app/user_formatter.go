package app

import "github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/models"

type UpdateUserFormatter struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type RegisterUserFormatter struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func RegisterFormatUser(user models.User, token string) RegisterUserFormatter {
	formatter := RegisterUserFormatter{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}
	return formatter
}

func GetFormatUser(user models.User) UpdateUserFormatter {
	formatter := UpdateUserFormatter{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
	return formatter
}
