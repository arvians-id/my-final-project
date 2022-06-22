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
	FindByUserCourse(ctx context.Context, code1 string, code2 string) (model.GetUserCourseResponse, error)
	Create(ctx context.Context, request model.CreateUserCourseRequest) (model.GetUserCourseResponse, error)
	Delete(ctx context.Context, code1 int, code2 int) error
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

func (service *usercourseService) FindByUserCourse(ctx context.Context, code1 string, code2 string) (model.GetUserCourseResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetUserCourseResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	usercourse, err := service.UserCourseRepository.FindByUserCourse(ctx, tx, code1, code2)
	if err != nil {
		return model.GetUserCourseResponse{}, err
	}

	return utils.ToUserCourseResponse(usercourse), nil
}

func (service *usercourseService) Create(ctx context.Context, request model.CreateUserCourseRequest) (model.GetUserCourseResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetUserCourseResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	usercourses := entity.UserCourse{
		UserId:   request.UserId,
		CourseId: request.CourseId,
	}

	usercourse, err := service.UserCourseRepository.Create(ctx, tx, usercourses)
	if err != nil {
		return model.GetUserCourseResponse{}, err
	}

	return utils.ToUserCourseResponse(usercourse), nil
}

func (service *usercourseService) Delete(ctx context.Context, code1 int, code2 int) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer utils.CommitOrRollback(tx)

	getCourse, err := service.UserCourseRepository.FindByUserCourse(ctx, tx, utils.ToString(code1), utils.ToString(code2))
	if err != nil {
		return err
	}

	err = service.UserCourseRepository.Delete(ctx, tx, getCourse.UserId, code2)
	if err != nil {
		return err
	}

	return nil
}
