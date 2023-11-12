package app

import "github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/models"

type RegisterUserInput struct {
	Username        string `json:"username" binding:"required" validate:"required,max=225,min=1"`
	Email           string `json:"email" binding:"required,email" validate:"required,max=225,min=1"`
	Password        string `json:"password" binding:"required" validate:"required,max=225,min=6"`
	ConfirmPassword string `json:"confirm_password" binding:"required" validate:"required,min=6,eqfield=Password"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email" validate:"required,min=1"`
	Password string `json:"password" binding:"required" validate:"required,max=225,min=6"`
}

type FormUpdateUserInput struct {
	Username        string `form:"username" binding:"required" validate:"required,max=225,min=1"`
	Email           string `form:"email" binding:"required,email" validate:"required,max=225,min=1"`
	Password        string `form:"password" binding:"required" validate:"required,max=225,min=6"`
	ConfirmPassword string `form:"confirm_password" binding:"required" validate:"required,max=225,min=6,eqfield=Password"`
	User            models.User
}
