package repositories

import (
	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/models"
	"gorm.io/gorm"
)

type PhotoRepositoryImpl struct {
	db *gorm.DB
}

func (r *PhotoRepositoryImpl) Save(photo models.Photo) (models.Photo, error) {
	err := r.db.Create(&photo).Error
	if err != nil {
		return photo, err
	}
	return photo, nil
}

func (r *PhotoRepositoryImpl) FindByID(ID int) (models.Photo, error) {
	var photo models.Photo
	err := r.db.Where("ID = ?", ID).Find(&photo).Error
	if err != nil {
		return photo, err
	}
	return photo, nil
}
func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &PhotoRepositoryImpl{db: db}
}
