package repositories

import (
	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/models"
)

type UserRepository interface {
	Save(user models.User) (models.User, error)
}
