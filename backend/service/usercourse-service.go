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
	FindAllStudentSubmissions(ctx context.Context, userId int, limit int) ([]model.GetStudentSubmissionsResponse, error)
	FindAllTeacherSubmissions(ctx context.Context, codeCourse string, moduleSubmissionId int) ([]model.GetTeacherSubmissionsResponse, error)
}

type usercourseService struct {
	UserCourseRepository       repository.UserCourseRepository
	CourseRepository           repository.CourseRepository
	ModuleSubmissionRepository repository.ModuleSubmissionsRepository
	DB                         *sql.DB
}

func NewUserCourseService(usercourseRepository *repository.UserCourseRepository, courseRepository *repository.CourseRepository, moduleSubmissionRepository *repository.ModuleSubmissionsRepository, db *sql.DB) UserCourseService {
	return &usercourseService{
		UserCourseRepository:       *usercourseRepository,
		CourseRepository:           *courseRepository,
		ModuleSubmissionRepository: *moduleSubmissionRepository,
		DB:                         db,
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

	array, err := service.UserCourseRepository.FindAll(ctx, tx)
	if err != nil {
		return model.GetUserCourseResponse{}, err
	}
	for _, usercourse := range array {
		if usercourse.UserId == usercourses.UserId && usercourse.CourseId == usercourses.CourseId {
			return model.GetUserCourseResponse{}, err
		}
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

func (service *usercourseService) FindAllStudentSubmissions(ctx context.Context, userId int, limit int) ([]model.GetStudentSubmissionsResponse, error) {

	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer utils.CommitOrRollback(tx)

	studentSubmissions, err := service.UserCourseRepository.FindAllStudentSubmissions(ctx, tx, userId, limit)
	if err != nil {
		return nil, err
	}

	var studentSubmissionsResponses []model.GetStudentSubmissionsResponse
	for _, studentSubmission := range studentSubmissions {
		studentSubmissionsResponses = append(studentSubmissionsResponses, utils.ToStudentSubmissionsResponse(studentSubmission))
	}

	return studentSubmissionsResponses, nil
}

func (service *usercourseService) FindAllTeacherSubmissions(ctx context.Context, codeCourse string, moduleSubmissionId int) ([]model.GetTeacherSubmissionsResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer utils.CommitOrRollback(tx)

	course, err := service.CourseRepository.FindByCode(ctx, tx, codeCourse)
	if err != nil {
		return nil, err
	}

	_, err = service.ModuleSubmissionRepository.FindByModId(ctx, tx, course.Id, moduleSubmissionId)
	if err != nil {
		return nil, err
	}

	teacherSubmissions, err := service.UserCourseRepository.FindAllTeacherSubmissions(ctx, tx, course.Id, moduleSubmissionId)
	if err != nil {
		return nil, err
	}

	var teacherSubmissionsResponses []model.GetTeacherSubmissionsResponse
	for _, studentSubmission := range teacherSubmissions {
		teacherSubmissionsResponses = append(teacherSubmissionsResponses, utils.ToTeacherSubmissionsResponse(studentSubmission))
	}

	return teacherSubmissionsResponses, nil
}
