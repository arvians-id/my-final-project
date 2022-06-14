package service

import (
	"context"
	"database/sql"
	"github.com/rg-km/final-project-engineering-12/backend/entity"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/repository"
	"github.com/rg-km/final-project-engineering-12/backend/utils"
)

type CourseService interface {
	FindAll(ctx context.Context, status bool, limit int) ([]model.GetCourseResponse, error)
	FindByCode(ctx context.Context, code string) (model.GetCourseResponse, error)
	Create(ctx context.Context, request model.CreateCourseRequest) (model.GetCourseResponse, error)
	Update(ctx context.Context, request model.UpdateCourseRequest, code string) (model.GetCourseResponse, error)
	Delete(ctx context.Context, code string) error
	ChangeActiveCourse(ctx context.Context, request model.UpdateStatusCourseRequest, code string) error
}

type courseService struct {
	CourseRepository repository.CourseRepository
	DB               *sql.DB
}

func NewCourseService(courseRepository *repository.CourseRepository, db *sql.DB) CourseService {
	return &courseService{
		CourseRepository: *courseRepository,
		DB:               db,
	}
}

func (service *courseService) FindAll(ctx context.Context, status bool, limit int) ([]model.GetCourseResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return []model.GetCourseResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	courses, err := service.CourseRepository.FindAll(ctx, tx, status, limit)
	if err != nil {
		return []model.GetCourseResponse{}, err
	}

	var courseResponses []model.GetCourseResponse
	for _, course := range courses {
		courseResponses = append(courseResponses, utils.ToCourseResponse(course))
	}

	return courseResponses, nil
}

func (service *courseService) FindByCode(ctx context.Context, code string) (model.GetCourseResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetCourseResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	course, err := service.CourseRepository.FindByCode(ctx, tx, code)
	if err != nil {
		return model.GetCourseResponse{}, err
	}

	return utils.ToCourseResponse(course), nil
}

func (service *courseService) Create(ctx context.Context, request model.CreateCourseRequest) (model.GetCourseResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetCourseResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	newCourse := entity.Courses{
		Name:        request.Name,
		CodeCourse:  utils.RandomString(10),
		Class:       request.Class,
		Tools:       request.Tools,
		About:       request.About,
		Description: request.Description,
		CreatedAt:   utils.TimeNow(),
		UpdatedAt:   utils.TimeNow(),
		IsActive:    true,
	}

	course, err := service.CourseRepository.Create(ctx, tx, newCourse)
	if err != nil {
		return model.GetCourseResponse{}, err
	}

	return utils.ToCourseResponse(course), nil
}

func (service *courseService) Update(ctx context.Context, request model.UpdateCourseRequest, code string) (model.GetCourseResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetCourseResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	getCourse, err := service.CourseRepository.FindByCode(ctx, tx, code)
	if err != nil {
		return model.GetCourseResponse{}, err
	}

	newCourse := entity.Courses{
		Name:        request.Name,
		CodeCourse:  code,
		Class:       request.Class,
		Tools:       request.Tools,
		About:       request.About,
		Description: request.Description,
		CreatedAt:   getCourse.CreatedAt,
		UpdatedAt:   utils.TimeNow(),
		IsActive:    getCourse.IsActive,
	}

	course, err := service.CourseRepository.Update(ctx, tx, newCourse, code)
	if err != nil {
		return model.GetCourseResponse{}, err
	}

	return utils.ToCourseResponse(course), nil
}

func (service *courseService) Delete(ctx context.Context, code string) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer utils.CommitOrRollback(tx)

	_, err = service.CourseRepository.FindByCode(ctx, tx, code)
	if err != nil {
		return err
	}

	err = service.CourseRepository.Delete(ctx, tx, code)
	if err != nil {
		return err
	}

	return nil
}

func (service *courseService) ChangeActiveCourse(ctx context.Context, request model.UpdateStatusCourseRequest, code string) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer utils.CommitOrRollback(tx)

	_, err = service.CourseRepository.FindByCode(ctx, tx, code)
	if err != nil {
		return err
	}

	err = service.CourseRepository.ChangeActiveCourse(ctx, tx, request.IsActive, code)
	if err != nil {
		return err
	}

	return nil
}
