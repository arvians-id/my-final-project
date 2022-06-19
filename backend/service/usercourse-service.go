package service

import (
	"context"
	"database/sql"

	"github.com/rg-km/final-project-engineering-12/backend/entity"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/repository"
	"github.com/rg-km/final-project-engineering-12/backend/utils"
)

type UserCourseService interface {
	FindAll(ctx context.Context) ([]model.GetUserCourseResponse, error)
	FindByUserId(ctx context.Context, code string) (model.GetUserCourseResponse, error)
	Create(ctx context.Context, requestusercourse model.CreateUserCourseRequest, requestusers model.UserDetailResponse, requestcourse model.GetCourseResponse) (model.GetUserCourseResponse, error)
	Delete(ctx context.Context, code string) error
}

type usercourseService struct {
	UserCourseRepository repository.UserCourseRepository
	DB                   *sql.DB
}

func NewUserCourseService(usercourseRepository *repository.UserCourseRepository, db *sql.DB) UserCourseService {
	return &usercourseService{
		UserCourseRepository: *usercourseRepository,
		DB:                   db,
	}
}

func (service *usercourseService) FindAll(ctx context.Context) ([]model.GetUserCourseResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return []model.GetUserCourseResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	courses, err := service.UserCourseRepository.FindAll(ctx, tx)
	if err != nil {
		return []model.GetUserCourseResponse{}, err
	}

	var usercourseResponses []model.GetUserCourseResponse
	for _, usercourse := range courses {
		usercourseResponses = append(usercourseResponses, utils.ToUserCourseResponse(usercourse))
	}

	return usercourseResponses, nil
}

func (service *usercourseService) FindByUserId(ctx context.Context, code string) (model.GetUserCourseResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetUserCourseResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	usercourse, err := service.UserCourseRepository.FindByUserId(ctx, tx, code)
	if err != nil {
		return model.GetUserCourseResponse{}, err
	}

	return utils.ToUserCourseResponse(usercourse), nil
}

func (service *usercourseService) Create(ctx context.Context, requestusercourse model.CreateUserCourseRequest, requestusers model.UserDetailResponse, requestcourse model.GetCourseResponse) (model.GetUserCourseResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetUserCourseResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	newUserCourse := entity.UserCourse{
		UserId:   requestusercourse.UserId,
		CourseId: requestusercourse.CourseId,
	}

	users := entity.Users{
		Id:       requestusers.Id,
		Username: requestusers.Username,
	}

	course := entity.Courses{
		Id:         requestcourse.Id,
		CodeCourse: requestcourse.CodeCourse,
	}

	usercourse, err := service.UserCourseRepository.Create(ctx, tx, newUserCourse, users, course)
	if err != nil {
		return model.GetUserCourseResponse{}, err
	}

	return utils.ToUserCourseResponse(usercourse), nil
}

func (service *usercourseService) Delete(ctx context.Context, code string) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer utils.CommitOrRollback(tx)

	getCourse, err := service.UserCourseRepository.FindByUserId(ctx, tx, code)
	if err != nil {
		return err
	}

	err = service.UserCourseRepository.Delete(ctx, tx, utils.ToString(getCourse.UserId))
	if err != nil {
		return err
	}

	return nil
}
