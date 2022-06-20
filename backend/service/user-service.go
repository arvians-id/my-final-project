package service

import (
<<<<<<< HEAD
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
	"fmt"
=======
	"context"
	"database/sql"
>>>>>>> Stashed changes
>>>>>>> Stashed changes
=======
	"context"
	"database/sql"
>>>>>>> 6ca9fa7d7d3ad5fb18980dbb0f7d514ea1b3a885
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-12/backend/entity"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/repository"
	"github.com/rg-km/final-project-engineering-12/backend/utils"
)

type UserService interface {
<<<<<<< HEAD
<<<<<<< Updated upstream
	RegisterUser(user model.UserRegisterResponse) (model.UserRegisterResponse, error)
	UserLogin(email string, password string) (model.UserLoginResponse, error)
	ListUser() ([]model.UserDetailResponse, error)
	GetUserbyID(id int) (model.UserDetailResponse, error)
	UpdateUser(id int, user model.UserDetailResponse) (model.UserDetailResponse, error)
=======
<<<<<<< Updated upstream
	RegisterUser(user model.UserRegister) (model.UserRegister, error)
	UserLogin(email string, password string) (model.UserRegister, error)
	ListUser() ([]model.UserRegister, error)
	GetUserbyID(id int) (model.UserRegister, error)
	UpdateUser(id int, user model.UserRegister) (model.UserRegister, error)
>>>>>>> Stashed changes
	DeleteUser(id int) error
=======
=======
>>>>>>> 6ca9fa7d7d3ad5fb18980dbb0f7d514ea1b3a885
	RegisterUser(ctx context.Context, user model.UserRegisterResponse) (model.UserRegisterResponse, error)
	UserLogin(ctx context.Context, user model.GetUserLogin) (model.UserLoginResponse, error)
	UpdateUserRole(ctx context.Context, id int) (model.UserDetailResponse, error)
	ListUser(ctx context.Context) ([]model.UserDetailResponse, error)
	GetUserbyID(ctx context.Context, id int) (model.UserDetailResponse, error)
	UpdateUser(ctx context.Context, id int, user model.UserDetailResponse) (model.UserDetailResponse, error)
	DeleteUser(ctx context.Context, id int) error
<<<<<<< HEAD
>>>>>>> Stashed changes
=======
>>>>>>> 6ca9fa7d7d3ad5fb18980dbb0f7d514ea1b3a885
}

type UserServiceImplement struct {
	userRepository repository.UserRepository
	DB             *sql.DB
}

func NewUserService(userRepository *repository.UserRepository, db *sql.DB) UserServiceImplement {
	return UserServiceImplement{
		userRepository: *userRepository,
		DB:             db,
	}
}

// RegisterUser is used to register new user
<<<<<<< HEAD
<<<<<<< Updated upstream
func (service *UserServiceImplement) RegisterUser(user model.UserRegisterResponse) (model.UserRegisterResponse, error) {
=======
func (service *UserServiceImplement) RegisterUser(ctx *gin.Context, user model.UserRegisterResponse) (model.UserRegisterResponse, error) {
>>>>>>> 6ca9fa7d7d3ad5fb18980dbb0f7d514ea1b3a885
	var response model.UserRegisterResponse
=======
<<<<<<< Updated upstream
func (service *UserServiceImplement) RegisterUser(user model.UserRegister) (model.UserRegister, error) {
	var response model.UserRegister
=======
func (service *UserServiceImplement) RegisterUser(ctx *gin.Context, user model.UserRegisterResponse) (model.UserRegisterResponse, error) {
	var response model.UserRegisterResponse
>>>>>>> Stashed changes

	tx, err := service.DB.Begin()

	if err != nil {
		return model.UserRegisterResponse{}, err
	}
	defer utils.CommitOrRollback(tx)
>>>>>>> Stashed changes

	tx, err := service.DB.Begin()

	if err != nil {
		return model.UserRegisterResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	user.Created_at = time.Now()
	user.Updated_at = time.Now()
	user.EmailVerification = time.Now()

	service.userRepository.Register(ctx, tx, entity.Users{
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

<<<<<<< HEAD
<<<<<<< Updated upstream
	temp, err := service.userRepository.GetLastInsertUser()
=======
<<<<<<< Updated upstream
	temp := service.userRepository.GetLastInsertUser()
	response = model.UserRegister{
=======
	temp, err := service.userRepository.GetLastInsertUser(ctx, tx)
>>>>>>> Stashed changes
=======
	temp, err := service.userRepository.GetLastInsertUser(ctx, tx)
>>>>>>> 6ca9fa7d7d3ad5fb18980dbb0f7d514ea1b3a885

	if err != nil {
		return model.UserRegisterResponse{}, err
	}

	response = model.UserRegisterResponse{
<<<<<<< Updated upstream
=======
>>>>>>> Stashed changes
>>>>>>> Stashed changes
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
<<<<<<< HEAD
<<<<<<< Updated upstream
func (service *UserServiceImplement) UserLogin(email string, password string) (model.UserLoginResponse, error) {
	var response model.UserLoginResponse
	user, err := service.userRepository.Login(email, password)
=======
<<<<<<< Updated upstream
func (service *UserServiceImplement) UserLogin(email string, password string) (model.UserRegister, error) {
	var response model.UserRegister
	user := service.userRepository.Login(email, password)
=======
func (service *UserServiceImplement) UserLogin(ctx *gin.Context, data model.GetUserLogin) (model.UserLoginResponse, error) {
	var response model.UserLoginResponse
=======
func (service *UserServiceImplement) UserLogin(ctx *gin.Context, data model.GetUserLogin) (model.UserLoginResponse, error) {
	var response model.UserLoginResponse
>>>>>>> 6ca9fa7d7d3ad5fb18980dbb0f7d514ea1b3a885

	tx, err := service.DB.Begin()

	if err != nil {
		return model.UserLoginResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	user, err := service.userRepository.Login(ctx, tx, data)
<<<<<<< HEAD

	if err != nil {
		return model.UserLoginResponse{}, err
	}
>>>>>>> Stashed changes
>>>>>>> Stashed changes
=======
>>>>>>> 6ca9fa7d7d3ad5fb18980dbb0f7d514ea1b3a885

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

<<<<<<< HEAD
<<<<<<< Updated upstream
func (service *UserServiceImplement) GetUserbyID(id int) (model.UserDetailResponse, error) {
	var response model.UserDetailResponse
	user, err := service.userRepository.GetUserByID(id)
=======
<<<<<<< Updated upstream
func (service *UserServiceImplement) GetUserbyID(id int) (model.UserRegister, error) {
	var response model.UserRegister
	user := service.userRepository.GetUserByID(id)
=======
// UpdateUserRole is used to update user role
func (service *UserServiceImplement) UpdateUserRole(ctx context.Context, id int) (model.UserDetailResponse, error) {
	var response model.UserDetailResponse
=======
// UpdateUserRole is used to update user role
func (service *UserServiceImplement) UpdateUserRole(ctx context.Context, id int) (model.UserDetailResponse, error) {
	var response model.UserDetailResponse
>>>>>>> 6ca9fa7d7d3ad5fb18980dbb0f7d514ea1b3a885

	tx, err := service.DB.Begin()

	if err != nil {
		return model.UserDetailResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	user, err := service.userRepository.UpdateRole(ctx, tx, id)

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

// GetUserbyID is used to get user by id
func (service *UserServiceImplement) GetUserbyID(ctx *gin.Context, id int) (model.UserDetailResponse, error) {
	var response model.UserDetailResponse

	tx, err := service.DB.Begin()

	if err != nil {
		return model.UserDetailResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	user, err := service.userRepository.GetUserByID(ctx, tx, id)
<<<<<<< HEAD

	if err != nil {
		return model.UserDetailResponse{}, err
	}
>>>>>>> Stashed changes
>>>>>>> Stashed changes
=======
>>>>>>> 6ca9fa7d7d3ad5fb18980dbb0f7d514ea1b3a885

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
<<<<<<< HEAD
<<<<<<< Updated upstream
func (service *UserServiceImplement) UpdateUser(id int, user model.GetUserDetailUpdate) (model.UserDetailResponse, error) {
=======
func (service *UserServiceImplement) UpdateUser(ctx *gin.Context, id int, user model.GetUserDetailUpdate) (model.UserDetailResponse, error) {
>>>>>>> 6ca9fa7d7d3ad5fb18980dbb0f7d514ea1b3a885
	var response model.UserDetailResponse
=======
<<<<<<< Updated upstream
func (service *UserServiceImplement) UpdateUser(id int, user model.UserRegister) (model.UserRegister, error) {
	var response model.UserRegister
>>>>>>> Stashed changes

	tx, err := service.DB.Begin()

	if err != nil {
		return model.UserDetailResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	user.UpdateAt = time.Now()

<<<<<<< HEAD
	err := service.userRepository.Update(entity.Users{
<<<<<<< Updated upstream
=======
		Id:                id,
		Name:              user.Name,
		Username:          user.Username,
		Email:             user.Email,
		Password:          user.Password,
		Role:              user.Role,
		EmailVerification: user.EmailVerification,
		CreatedAt:         user.Created_at,
		UpdatedAt:         user.Updated_at,
=======
func (service *UserServiceImplement) UpdateUser(ctx *gin.Context, id int, user model.GetUserDetailUpdate) (model.UserDetailResponse, error) {
	var response model.UserDetailResponse

	tx, err := service.DB.Begin()

	if err != nil {
		return model.UserDetailResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	user.UpdateAt = time.Now()

	err = service.userRepository.Update(ctx, tx, entity.Users{
>>>>>>> Stashed changes
=======
	err = service.userRepository.Update(ctx, tx, entity.Users{
>>>>>>> 6ca9fa7d7d3ad5fb18980dbb0f7d514ea1b3a885
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
<<<<<<< Updated upstream
=======
>>>>>>> Stashed changes
>>>>>>> Stashed changes
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
<<<<<<< HEAD
<<<<<<< Updated upstream
func (service *UserServiceImplement) ListUser() ([]model.UserDetailResponse, error) {
	var responses = []model.UserDetailResponse{}
	users, err := service.userRepository.ListUser()
=======
<<<<<<< Updated upstream
func (service *UserServiceImplement) ListUser() ([]model.UserRegister, error) {
	var responses = []model.UserRegister{}
	users := service.userRepository.ListUser()
=======
func (service *UserServiceImplement) ListUser(ctx *gin.Context) ([]model.UserDetailResponse, error) {
	var responses = []model.UserDetailResponse{}
=======
func (service *UserServiceImplement) ListUser(ctx *gin.Context) ([]model.UserDetailResponse, error) {
	var responses = []model.UserDetailResponse{}
>>>>>>> 6ca9fa7d7d3ad5fb18980dbb0f7d514ea1b3a885

	tx, err := service.DB.Begin()

	if err != nil {
		return []model.UserDetailResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	users, err := service.userRepository.ListUser(ctx, tx)
<<<<<<< HEAD
>>>>>>> Stashed changes
=======
>>>>>>> 6ca9fa7d7d3ad5fb18980dbb0f7d514ea1b3a885

	if err != nil {
		return nil, err
	}
<<<<<<< Updated upstream
=======
>>>>>>> Stashed changes
>>>>>>> Stashed changes

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
func (service *UserServiceImplement) DeleteUser(ctx *gin.Context, id int) error {

	tx, err := service.DB.Begin()

	if err != nil {
		return err
	}
	defer utils.CommitOrRollback(tx)

	err = service.userRepository.Delete(ctx, tx, id)

	if err != nil {
		return err
	}

	return nil
}
