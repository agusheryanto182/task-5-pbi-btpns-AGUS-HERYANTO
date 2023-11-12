package services

import (
	"errors"

	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/app"
	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/models"
	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/repositories"
	"github.com/go-playground/validator/v10"
)

type PhotoServiceImpl struct {
	photoRepository repositories.PhotoRepository
	validate        *validator.Validate
}

func (s *PhotoServiceImpl) Create(inputID int, inputData app.PhotoInput) (models.Photo, error) {
	// err := s.validate.Struct(inputData)
	// if err != nil {
	// 	return models.Photo{}, err
	// }

	photo := models.Photo{}
	photo.Title = inputData.Title
	photo.Caption = inputData.Caption
	photo.PhotoURL = inputData.PhotoURL
	photo.UserID = inputID

	created, err := s.photoRepository.Save(photo)
	if err != nil {
		return created, err
	}
	return created, nil
}

func (s *PhotoServiceImpl) GetByUserID(userID int) ([]models.Photo, error) {
	photo, err := s.photoRepository.FindByUserID(userID)
	if err != nil {
		return photo, err
	}
	return photo, nil
}

func (s *PhotoServiceImpl) GetByID(ID int) (models.Photo, error) {
	photo, err := s.photoRepository.FindByID(ID)
	if err != nil {
		return photo, err
	}

	if photo.ID == 0 {
		return photo, errors.New("Photo not found")
	}
	return photo, nil
}

func (s *PhotoServiceImpl) Update(inputID int, inputData app.PhotoUpdate) (models.Photo, error) {
	// err := s.validate.Struct(inputData)
	// if err != nil {
	// 	return models.Photo{}, err
	// }

	photo, err := s.photoRepository.FindByID(inputID)
	if err != nil {
		return photo, err
	}

	photo.Title = inputData.Title
	photo.Caption = inputData.Caption

	updatedPhoto, err := s.photoRepository.Update(photo)
	if err != nil {
		return updatedPhoto, err
	}

	return updatedPhoto, nil

}

func (s *PhotoServiceImpl) Delete(ID int) error {
	err := s.photoRepository.Delete(ID)
	if err != nil {
		return err
	}
	return nil
}

func NewPhotoService(photoRepository repositories.PhotoRepository, validate *validator.Validate) PhotoService {
	return &PhotoServiceImpl{photoRepository: photoRepository, validate: validate}
}
