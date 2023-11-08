package repositories

import "github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/models"

type PhotoRepository interface {
	Save(photo models.Photo) (models.Photo, error)
	FindByID(ID int) (models.Photo, error)
	Update(photo models.Photo) (models.Photo, error)
	Delete(ID int) error
}
