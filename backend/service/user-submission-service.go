package service

import (
	"context"
	"database/sql"
	"github.com/rg-km/final-project-engineering-12/backend/entity"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/repository"
	"github.com/rg-km/final-project-engineering-12/backend/utils"
	"os"
)

type UserSubmissionsService interface {
	SubmitFile(ctx context.Context, request model.CreateUserSubmissionsRequest) (model.GetUserSubmissionsResponse, error)
	UpdateGrade(ctx context.Context, request model.UpdateUserGradeRequest) error
	FindUserSubmissionById(ctx context.Context, id int, userId int, moduleSubmissionsId int) (model.GetUserSubmissionsResponse, error)
}

type userSubmissionsService struct {
	UserSubmissionRepository    repository.UserSubmissionsRepository
	ModuleSubmissionsRepository repository.ModuleSubmissionsRepository
	CourseRepository            repository.CourseRepository
	DB                          *sql.DB
}

func NewUserSubmissionsService(userSubmissionRepository *repository.UserSubmissionsRepository, moduleSubmissionsRepository *repository.ModuleSubmissionsRepository, courseRepository *repository.CourseRepository, db *sql.DB) UserSubmissionsService {
	return &userSubmissionsService{
		UserSubmissionRepository:    *userSubmissionRepository,
		ModuleSubmissionsRepository: *moduleSubmissionsRepository,
		CourseRepository:            *courseRepository,
		DB:                          db,
	}
}

func (service *userSubmissionsService) FindUserSubmissionById(ctx context.Context, id int, userId int, moduleSubmissionsId int) (model.GetUserSubmissionsResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetUserSubmissionsResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	getUserSubmission := entity.UserSubmissions{
		Id:                 id,
		UserId:             userId,
		ModuleSubmissionId: moduleSubmissionsId,
	}
	userSubmission, err := service.UserSubmissionRepository.FindUserSubmissionById(ctx, tx, getUserSubmission)
	if err != nil {
		return model.GetUserSubmissionsResponse{}, err
	}

	return utils.ToUserSubmissionsResponse(userSubmission), nil
}

func (service *userSubmissionsService) SubmitFile(ctx context.Context, request model.CreateUserSubmissionsRequest) (model.GetUserSubmissionsResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetUserSubmissionsResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	newSubmit := entity.UserSubmissions{
		UserId:             request.UserId,
		ModuleSubmissionId: request.ModuleSubmissionId,
		File:               request.File,
	}

	before, err := service.UserSubmissionRepository.FindUserSubmissionById(ctx, tx, newSubmit)
	if err != nil {
		userSubmission, err := service.UserSubmissionRepository.SubmitFile(ctx, tx, newSubmit)
		if err != nil {
			return model.GetUserSubmissionsResponse{}, err
		}

		return utils.ToUserSubmissionsResponse(userSubmission), nil
	}
	userSubmission, err := service.UserSubmissionRepository.UpdateFile(ctx, tx, newSubmit)
	if err != nil {
		return model.GetUserSubmissionsResponse{}, err
	}

	path, err := utils.GetPath("/assets/", before.File)
	if err != nil {
		return model.GetUserSubmissionsResponse{}, err
	}

	err = os.Remove(path)
	if err != nil {
		return model.GetUserSubmissionsResponse{}, err
	}

	return utils.ToUserSubmissionsResponse(userSubmission), nil
}

func (service *userSubmissionsService) UpdateGrade(ctx context.Context, request model.UpdateUserGradeRequest) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer utils.CommitOrRollback(tx)

	newUpdate := entity.UserSubmissions{
		Id:    request.Id,
		Grade: &request.Grade,
	}

	err = service.UserSubmissionRepository.UpdateGrade(ctx, tx, newUpdate)
	if err != nil {
		return err
	}

	return nil
}
