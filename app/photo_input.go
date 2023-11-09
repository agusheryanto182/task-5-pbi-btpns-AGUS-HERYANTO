package app

import "github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/models"

type PhotoInput struct {
	Title    string      `json:"title" binding:"required" validate:"required,max=225,min=1"`
	Caption  string      `json:"caption" binding:"required" validate:"required,max=225,min=1"`
	PhotoURL string      `json:"photo_url" binding:"required" validate:"required,max=225,min=1"`
	UserID   models.User `json:"user_id" binding:"required" validate:"required,max=225,min=1"`
}
