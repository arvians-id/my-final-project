package service

import (
	"net/http"

	"github.com/rg-km/final-project-engineering-12/backend/model"
)

type UserService interface {
	GetAllUser() ([]model.UserRegister, error)
	RegisterUser(user model.UserRegister) (model.UserRegister, error)
	GetUserbyID(id int) (model.UserRegister, error)
	DeleteUser(id int) error
	UpdateUser(id int, user model.UserRegister) (model.UserRegister, error)
	UserLogin(email string, password string) (model.UserRegister, error)
	UserAuth(cookie http.Cookie) (model.UserRegister, error)
}
