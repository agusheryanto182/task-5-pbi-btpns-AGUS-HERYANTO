package repositories

import (
	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/models"
)

type UserRepository interface {
	Save(user models.User) (models.User, error)
	FindByEmail(email string) (models.User, error)
	FindByID(ID int) (models.User, error)
	Update(user models.User) (models.User, error)
	Delete(ID int) error
}
