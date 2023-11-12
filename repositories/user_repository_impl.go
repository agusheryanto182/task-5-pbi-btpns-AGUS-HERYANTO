package repositories

import (
	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/models"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (r *UserRepositoryImpl) Save(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) FindByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) FindByUsername(username string) (models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) FindByID(ID int) (models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) Update(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) Delete(ID int) error {
	var user models.User
	err := r.db.Where("id = ?", ID).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}
