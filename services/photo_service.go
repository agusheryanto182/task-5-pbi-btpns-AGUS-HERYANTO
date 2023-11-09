package services

import (
	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/app"
	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/models"
)

type PhotoService interface {
	Create(input app.PhotoInput) (models.Photo, error)
	GetByID(ID int) (models.Photo, error)
	Update(inputID int, inputData app.PhotoInput) (models.Photo, error)
	Delete(ID int) error
}
