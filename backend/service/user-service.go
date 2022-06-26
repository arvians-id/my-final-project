package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-12/backend/entity"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/repository"
	"github.com/rg-km/final-project-engineering-12/backend/utils"
)

type UserService interface {
	RegisterUser(ctx context.Context, user model.UserRegisterResponse, signature string, expired int) (model.UserRegisterResponse, error)
	UserLogin(ctx context.Context, user model.GetUserLogin) (model.UserLoginResponse, error)
	UpdateUserRole(ctx context.Context, id int) (model.UserDetailResponse, error)
	ListUser(ctx context.Context) ([]model.UserDetailResponse, error)
	GetUserbyID(ctx context.Context, id int) (model.UserDetailResponse, error)
	UpdateUser(ctx context.Context, id int, user model.UserDetailResponse) (model.UserDetailResponse, error)
	DeleteUser(ctx context.Context, id int) error
}

type UserServiceImplement struct {
	userRepository    repository.UserRepository
	emailVerification repository.EmailVerificationRepository
	DB                *sql.DB
}

func NewUserService(userRepository *repository.UserRepository, db *sql.DB, emailVerification *repository.EmailVerificationRepository) UserServiceImplement {
	return UserServiceImplement{
		userRepository:    *userRepository,
		emailVerification: *emailVerification,
		DB:                db,
	}
}

// RegisterUser is used to register new user
func (service *UserServiceImplement) RegisterUser(ctx *gin.Context, user model.UserRegisterResponse, signature string, expired int) (model.UserRegisterResponse, error) {
	var response model.UserRegisterResponse

	tx, err := service.DB.Begin()

	if err != nil {
		return model.UserRegisterResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	user.Created_at = time.Now()
	user.Updated_at = time.Now()
	user.EmailVerification = time.Now()

	err = service.userRepository.Register(ctx, tx, entity.Users{
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

	if err != nil {
		return model.UserRegisterResponse{}, err
	}

	temp, err := service.userRepository.GetLastInsertUser(ctx, tx)
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

	// Send Email Verification
	// Check If email is exists in table email verifications
	rows, err := service.emailVerification.FindByEmail(ctx, tx, user.Email)
	if err != nil {
		return model.UserRegisterResponse{}, err
	}
	// Create new signature if not exist
	emailVerification := entity.EmailVerification{
		Email:     user.Email,
		Signature: signature,
		Expired:   expired,
	}
	if rows.Email == "" {
		_, err = service.emailVerification.Create(ctx, tx, emailVerification)
		if err != nil {
			return model.UserRegisterResponse{}, err
		}
	} else {
		// Update token if email is exist
		_, err = service.emailVerification.Update(ctx, tx, emailVerification)
		if err != nil {
			return model.UserRegisterResponse{}, err
		}
	}

	return response, nil
}

// UserLogin is used to login user
func (service *UserServiceImplement) UserLogin(ctx *gin.Context, data model.GetUserLogin) (model.UserLoginResponse, error) {
	var response model.UserLoginResponse

	tx, err := service.DB.Begin()

	if err != nil {
		return model.UserLoginResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	user, err := service.userRepository.Login(ctx, tx, data)

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

// UpdateUserRole is used to update user role
func (service *UserServiceImplement) UpdateUserRole(ctx context.Context, id int) (model.UserDetailResponse, error) {
	var response model.UserDetailResponse

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
func (service *UserServiceImplement) UpdateUser(ctx *gin.Context, id int, user model.GetUserDetailUpdate) (model.UserDetailResponse, error) {
	var response model.UserDetailResponse

	tx, err := service.DB.Begin()

	if err != nil {
		return model.UserDetailResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	user.UpdateAt = time.Now()

	err = service.userRepository.Update(ctx, tx, entity.Users{
		Id:             id,
		Name:           user.Name,
		Username:       user.Username,
		Role:           user.Role,
		Phone:          user.Phone,
		Gender:         user.Gender,
		DisabilityType: user.DisabilityType,
		Address:        &user.Address,
		Birthdate:      user.Birthdate,
		Image:          &user.Image,
		Description:    &user.Description,
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
		Address:        &user.Address,
		Birthdate:      user.Birthdate,
		Image:          &user.Image,
		Description:    &user.Description,
	}

	return response, nil
}

// ListUser is used to list all user
func (service *UserServiceImplement) ListUser(ctx *gin.Context) ([]model.UserDetailResponse, error) {
	var responses = []model.UserDetailResponse{}

	tx, err := service.DB.Begin()

	if err != nil {
		return []model.UserDetailResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	users, err := service.userRepository.ListUser(ctx, tx)

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
