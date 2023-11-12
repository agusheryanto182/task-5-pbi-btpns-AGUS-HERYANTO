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

func (r *PhotoRepositoryImpl) FindByUserID(userID int) ([]models.Photo, error) {
	var photo []models.Photo
	err := r.db.Where("user_id = ?", userID).Find(&photo).Error
	if err != nil {
		return photo, err
	}
	return photo, nil
}

func (r *PhotoRepositoryImpl) FindByID(ID int) (models.Photo, error) {
	var photo models.Photo
	err := r.db.Where("id = ?", ID).Find(&photo).Error
	if err != nil {
		return photo, err
	}
	return photo, nil
}

func (r *PhotoRepositoryImpl) Update(photo models.Photo) (models.Photo, error) {
	err := r.db.Save(&photo).Error
	if err != nil {
		return photo, err
	}
	return photo, nil
}

func (r *PhotoRepositoryImpl) Delete(ID int) error {
	var photo models.Photo
	err := r.db.Where("id = ?", ID).Delete(&photo).Error
	if err != nil {
		return err
	}
	return nil
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &PhotoRepositoryImpl{db: db}
}
