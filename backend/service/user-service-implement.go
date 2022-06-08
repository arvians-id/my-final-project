package service

import (
	"time"

	"github.com/rg-km/final-project-engineering-12/backend/entity"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/repository"
)

func NewUserService(userRepository *repository.UserRepository) UserServiceImplement {
	return UserServiceImplement{
		userRepository: *userRepository,
	}
}

type UserServiceImplement struct {
	userRepository repository.UserRepository
}

func (service *UserServiceImplement) GetAllUser() ([]model.UserRegister, error) {
	var responses = []model.UserRegister{}
	users := service.userRepository.FindUser()

	for _, user := range users {
		responses = append(responses, model.UserRegister{
			ID:            user.ID,
			Name:          user.Name,
			Unique_number: user.Unique_number,
			Phone:         user.Phone,
			Email:         user.Email,
			Password:      user.Password,
			Role:          user.Role,
			Image:         user.Image,
			Created_at:    user.Created_at,
			Updated_at:    user.Updated_at,
		})
	}

	return responses, nil
}

func (service *UserServiceImplement) RegisterUser(user model.UserRegister) (model.UserRegister, error) {
	var response model.UserRegister

	user.Created_at = time.Now()
	user.Updated_at = time.Now()

	service.userRepository.Insert(entity.User{
		ID:            user.ID,
		Name:          user.Name,
		Unique_number: user.Unique_number,
		Phone:         user.Phone,
		Email:         user.Email,
		Password:      user.Password,
		Role:          user.Role,
		Image:         user.Image,
		Created_at:    user.Created_at,
		Updated_at:    user.Updated_at,
	})

	response = user

	return response, nil
}
