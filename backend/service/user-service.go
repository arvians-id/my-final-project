package service

import (
	"time"

	"github.com/rg-km/final-project-engineering-12/backend/entity"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/repository"
)

type UserService interface {
	RegisterUser(user model.UserRegisterResponse) (model.UserRegisterResponse, error)
	UserLogin(email string, password string) (model.UserLoginResponse, error)
	ListUser() ([]model.UserDetailResponse, error)
	GetUserbyID(id int) (model.UserDetailResponse, error)
	UpdateUser(id int, user model.UserDetailResponse) (model.UserDetailResponse, error)
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
func (service *UserServiceImplement) RegisterUser(user model.UserRegisterResponse) (model.UserRegisterResponse, error) {
	var response model.UserRegisterResponse

	user.Created_at = time.Now()
	user.Updated_at = time.Now()
	user.EmailVerification = time.Now()

	service.userRepository.Register(entity.Users{
		Name:              user.Name,
		Username:          user.Username,
		Email:             user.Email,
		Password:          user.Password,
		Role:              user.Role,
		Phone:             user.Phone,
		Gender:            user.Gender,
		DisabilityType:    user.DisabilityType,
		Birthdate:         user.Birthdate,
		EmailVerification: user.EmailVerification,
		CreatedAt:         user.Created_at,
		UpdatedAt:         user.Updated_at,
	})

	temp, err := service.userRepository.GetLastInsertUser()

	if err != nil {
		return model.UserRegisterResponse{}, err
	}

	response = model.UserRegisterResponse{
		Id:                temp.Id,
		Name:              temp.Name,
		Username:          temp.Username,
		Email:             temp.Email,
		Password:          temp.Password,
		Role:              temp.Role,
		Phone:             temp.Phone,
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
func (service *UserServiceImplement) UserLogin(email string, password string) (model.UserLoginResponse, error) {
	var response model.UserLoginResponse
	user, err := service.userRepository.Login(email, password)

	if err != nil {
		return model.UserLoginResponse{}, err
	}

	response = model.UserLoginResponse{
		Id:             user.Id,
		Name:           user.Name,
		Username:       user.Username,
		Email:          user.Email,
		Role:           user.Role,
		Gender:         user.Gender,
		DisabilityType: user.DisabilityType,
	}
	return response, nil
}

func (service *UserServiceImplement) GetUserbyID(id int) (model.UserDetailResponse, error) {
	var response model.UserDetailResponse
	user, err := service.userRepository.GetUserByID(id)

	if err != nil {
		return model.UserDetailResponse{}, err
	}

	response = model.UserDetailResponse{
		Id:             user.Id,
		Name:           user.Name,
		Username:       user.Username,
		Role:           user.Role,
		Phone:          user.Phone,
		Gender:         user.Gender,
		DisabilityType: user.DisabilityType,
		Address:        user.Address,
		Birthdate:      user.Birthdate,
		Image:          user.Image,
		Description:    user.Description,
	}

	return response, nil
}

// UpdateUser is used to update user
func (service *UserServiceImplement) UpdateUser(id int, user model.GetUserDetailUpdate) (model.UserDetailResponse, error) {
	var response model.UserDetailResponse

	user.UpdateAt = time.Now()

	err := service.userRepository.Update(entity.Users{
		Id:             id,
		Name:           user.Name,
		Username:       user.Username,
		Role:           user.Role,
		Phone:          user.Phone,
		Gender:         user.Gender,
		DisabilityType: user.DisabilityType,
		Address:        user.Address,
		Birthdate:      user.Birthdate,
		Image:          user.Image,
		Description:    user.Description,
		UpdatedAt:      user.UpdateAt,
	})

	if err != nil {
		return model.UserDetailResponse{}, err
	}

	response = model.UserDetailResponse{
		Id:             id,
		Name:           user.Name,
		Username:       user.Username,
		Role:           user.Role,
		Phone:          user.Phone,
		Gender:         user.Gender,
		DisabilityType: user.DisabilityType,
		Address:        user.Address,
		Birthdate:      user.Birthdate,
		Image:          user.Image,
		Description:    user.Description,
	}

	return response, nil
}

// ListUser is used to list all user
func (service *UserServiceImplement) ListUser() ([]model.UserDetailResponse, error) {
	var responses = []model.UserDetailResponse{}
	users, err := service.userRepository.ListUser()

	if err != nil {
		return nil, err
	}

	for _, user := range users {
		responses = append(responses, model.UserDetailResponse{
			Id:             user.Id,
			Name:           user.Name,
			Username:       user.Username,
			Role:           user.Role,
			Phone:          user.Phone,
			Gender:         user.Gender,
			DisabilityType: user.DisabilityType,
			Address:        user.Address,
			Birthdate:      user.Birthdate,
			Image:          user.Image,
			Description:    user.Description,
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
