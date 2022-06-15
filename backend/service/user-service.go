package service

import (
	"fmt"
	"time"

	"github.com/rg-km/final-project-engineering-12/backend/entity"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/repository"
)

type UserService interface {
	RegisterUser(user model.UserRegister) (model.UserRegister, error)
	UserLogin(email string, password string) (model.UserRegister, error)
	ListUser() ([]model.UserRegister, error)
	GetUserbyID(id int) (model.UserRegister, error)
	UpdateUser(id int, user model.UserRegister) (model.UserRegister, error)
	DeleteUser(id int) error
}

func NewUserService(userRepository *repository.UserRepository) UserServiceImplement {
	return UserServiceImplement{
		userRepository: *userRepository,
	}
}

type UserServiceImplement struct {
	userRepository repository.UserRepository
}

// RegisterUser is used to register new user
func (service *UserServiceImplement) RegisterUser(user model.UserRegister) (model.UserRegister, error) {
	var response model.UserRegister

	user.Created_at = time.Now()
	user.Updated_at = time.Now()
	user.EmailVerification = time.Now()

	service.userRepository.Register(entity.Users{
		Name:              user.Name,
		Username:          user.Username,
		Email:             user.Email,
		Password:          user.Password,
		Role:              user.Role,
		Gender:            user.Gender,
		DisabilityType:    user.DisabilityType,
		Birthdate:         user.Birthdate,
		EmailVerification: user.EmailVerification,
		CreatedAt:         user.Created_at,
		UpdatedAt:         user.Updated_at,
	})

	temp := service.userRepository.GetLastInsertUser()
	response = model.UserRegister{
		Id:                temp.Id,
		Name:              temp.Name,
		Username:          temp.Username,
		Email:             temp.Email,
		Password:          temp.Password,
		Role:              temp.Role,
		Gender:            temp.Gender,
		DisabilityType:    user.DisabilityType,
		Birthdate:         user.Birthdate,
		EmailVerification: temp.EmailVerification,
		Created_at:        temp.CreatedAt,
		Updated_at:        temp.UpdatedAt,
	}

	return response, nil
}

// UserLogin is used to login user
func (service *UserServiceImplement) UserLogin(email string, password string) (model.UserRegister, error) {
	var response model.UserRegister
	user := service.userRepository.Login(email, password)

	response = model.UserRegister{
		Id:                user.Id,
		Name:              user.Name,
		Username:          user.Username,
		Email:             user.Email,
		Password:          user.Password,
		Role:              user.Role,
		Gender:            user.Gender,
		DisabilityType:    user.DisabilityType,
		Birthdate:         user.Birthdate,
		EmailVerification: user.EmailVerification,
		Created_at:        user.CreatedAt,
		Updated_at:        user.UpdatedAt,
	}
	return response, nil
}

func (service *UserServiceImplement) GetUserbyID(id int) (model.UserRegister, error) {
	var response model.UserRegister
	user := service.userRepository.GetUserByID(id)

	response = model.UserRegister{
		Id:                user.Id,
		Name:              user.Name,
		Username:          user.Username,
		Email:             user.Email,
		Password:          user.Password,
		Role:              user.Role,
		EmailVerification: user.EmailVerification,
		Created_at:        user.CreatedAt,
		Updated_at:        user.UpdatedAt,
	}

	return response, nil
}

// UpdateUser is used to update user
func (service *UserServiceImplement) UpdateUser(id int, user model.UserRegister) (model.UserRegister, error) {
	var response model.UserRegister

	user.Updated_at = time.Now()

	err := service.userRepository.Update(entity.Users{
		Id:                id,
		Name:              user.Name,
		Username:          user.Username,
		Email:             user.Email,
		Password:          user.Password,
		Role:              user.Role,
		EmailVerification: user.EmailVerification,
		CreatedAt:         user.Created_at,
		UpdatedAt:         user.Updated_at,
	})

	if err != nil {
		return model.UserRegister{}, fmt.Errorf(err.Error())
	}

	response = model.UserRegister(user)

	return response, nil
}

// ListUser is used to list all user
func (service *UserServiceImplement) ListUser() ([]model.UserRegister, error) {
	var responses = []model.UserRegister{}
	users := service.userRepository.ListUser()

	for _, user := range users {
		responses = append(responses, model.UserRegister{
			Id:                user.Id,
			Name:              user.Name,
			Username:          user.Username,
			Email:             user.Email,
			Password:          user.Password,
			Role:              user.Role,
			EmailVerification: user.EmailVerification,
			Created_at:        user.CreatedAt,
			Updated_at:        user.UpdatedAt,
		})
	}

	return responses, nil
}

// DeleteUser is used to delete user
func (service *UserServiceImplement) DeleteUser(id int) error {
	err := service.userRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
