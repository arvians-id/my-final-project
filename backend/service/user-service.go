package service

import (
	"github.com/rg-km/final-project-engineering-12/backend/model"
)

type UserService interface {
	GetAllUser() ([]model.UserRegister, error)
	RegisterUser(user model.UserRegister) (model.UserRegister, error)
}
