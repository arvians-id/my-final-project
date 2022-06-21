package service

import (
	"context"
	"database/sql"
	"github.com/rg-km/final-project-engineering-12/backend/entity"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/repository"
	"github.com/rg-km/final-project-engineering-12/backend/utils"
)

type ModuleSubmissionsService interface {
	FindAll(ctx context.Context, code string) ([]model.GetModuleSubmissionsResponse, error)
	FindByModId(ctx context.Context, code string, idArticle int) (model.GetModuleSubmissionsResponse, error)
	Create(ctx context.Context, request model.CreateModuleSubmissionsRequest, code string) (model.GetModuleSubmissionsResponse, error)
	Update(ctx context.Context, request model.UpdateModuleSubmissionsRequest, code string, idArticle int) (model.GetModuleSubmissionsResponse, error)
	Delete(ctx context.Context, code string, idArticle int) error
	Next(ctx context.Context, code string, idArticle int) (model.GetNextPreviousSubmissionsResponse, error)
	Previous(ctx context.Context, code string, idArticle int) (model.GetNextPreviousSubmissionsResponse, error)
}

type moduleSubmissionsService struct {
	ModuleSubmissionsRepository repository.ModuleSubmissionsRepository
	CourseRepository            repository.CourseRepository
	DB                          *sql.DB
}

func NewModuleSubmissionsService(moduleSubmissionsRepository *repository.ModuleSubmissionsRepository, courseRepository *repository.CourseRepository, db *sql.DB) ModuleSubmissionsService {
	return &moduleSubmissionsService{
		ModuleSubmissionsRepository: *moduleSubmissionsRepository,
		CourseRepository:            *courseRepository,
		DB:                          db,
	}
}

func (service *moduleSubmissionsService) FindAll(ctx context.Context, code string) ([]model.GetModuleSubmissionsResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return []model.GetModuleSubmissionsResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	course, err := service.CourseRepository.FindByCode(ctx, tx, code)
	if err != nil {
		return []model.GetModuleSubmissionsResponse{}, err
	}

	modsubs, err := service.ModuleSubmissionsRepository.FindAll(ctx, tx, course.Id)
	if err != nil {
		return []model.GetModuleSubmissionsResponse{}, err
	}

	var modsubResponses []model.GetModuleSubmissionsResponse
	for _, modsub := range modsubs {
		modsubResponses = append(modsubResponses, utils.ToModuleSubmissionsResponse(modsub))
	}

	return modsubResponses, nil
}

func (service *moduleSubmissionsService) FindByModId(ctx context.Context, code string, idArticle int) (model.GetModuleSubmissionsResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetModuleSubmissionsResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	course, err := service.CourseRepository.FindByCode(ctx, tx, code)
	if err != nil {
		return model.GetModuleSubmissionsResponse{}, err
	}

	modsub, err := service.ModuleSubmissionsRepository.FindByModId(ctx, tx, course.Id, idArticle)
	if err != nil {
		return model.GetModuleSubmissionsResponse{}, err
	}

	return utils.ToModuleSubmissionsResponse(modsub), nil
}

func (service *moduleSubmissionsService) Create(ctx context.Context, request model.CreateModuleSubmissionsRequest, code string) (model.GetModuleSubmissionsResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetModuleSubmissionsResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	course, err := service.CourseRepository.FindByCode(ctx, tx, code)
	if err != nil {
		return model.GetModuleSubmissionsResponse{}, err
	}

	newModsub := entity.ModuleSubmissions{
		CourseId:    course.Id,
		Name:        request.Name,
		Description: request.Description,
		Deadline:    utils.ParseTime(request.Deadline),
	}

	modsub, err := service.ModuleSubmissionsRepository.Create(ctx, tx, newModsub)
	if err != nil {
		return model.GetModuleSubmissionsResponse{}, err
	}

	return utils.ToModuleSubmissionsResponse(modsub), nil
}

func (service *moduleSubmissionsService) Update(ctx context.Context, request model.UpdateModuleSubmissionsRequest, code string, idArticle int) (model.GetModuleSubmissionsResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetModuleSubmissionsResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	course, err := service.CourseRepository.FindByCode(ctx, tx, code)
	if err != nil {
		return model.GetModuleSubmissionsResponse{}, err
	}

	_, err = service.ModuleSubmissionsRepository.FindByModId(ctx, tx, course.Id, idArticle)
	if err != nil {
		return model.GetModuleSubmissionsResponse{}, err
	}

	newModsub := entity.ModuleSubmissions{
		CourseId:    course.Id,
		Name:        request.Name,
		Description: request.Description,
		Deadline:    utils.ParseTime(request.Deadline),
	}

	modsub, err := service.ModuleSubmissionsRepository.Update(ctx, tx, newModsub, idArticle)
	if err != nil {
		return model.GetModuleSubmissionsResponse{}, err
	}

	return utils.ToModuleSubmissionsResponse(modsub), nil
}

func (service *moduleSubmissionsService) Delete(ctx context.Context, code string, idArticle int) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer utils.CommitOrRollback(tx)

	course, err := service.CourseRepository.FindByCode(ctx, tx, code)
	if err != nil {
		return err
	}

	_, err = service.ModuleSubmissionsRepository.FindByModId(ctx, tx, course.Id, idArticle)
	if err != nil {
		return err
	}

	err = service.ModuleSubmissionsRepository.Delete(ctx, tx, idArticle)
	if err != nil {
		return err
	}

	return nil
}

func (service *moduleSubmissionsService) Next(ctx context.Context, code string, idArticle int) (model.GetNextPreviousSubmissionsResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetNextPreviousSubmissionsResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	course, err := service.CourseRepository.FindByCode(ctx, tx, code)
	if err != nil {
		return model.GetNextPreviousSubmissionsResponse{}, err
	}

	_, err = service.ModuleSubmissionsRepository.FindByModId(ctx, tx, course.Id, idArticle)
	if err != nil {
		return model.GetNextPreviousSubmissionsResponse{}, err
	}

	next, err := service.ModuleSubmissionsRepository.Next(ctx, tx, course.Id, idArticle)
	if err != nil {
		return model.GetNextPreviousSubmissionsResponse{}, err
	}

	return utils.ToModuleSubmissionsNextPreviousResponse(next), nil
}

func (service *moduleSubmissionsService) Previous(ctx context.Context, code string, idArticle int) (model.GetNextPreviousSubmissionsResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetNextPreviousSubmissionsResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	course, err := service.CourseRepository.FindByCode(ctx, tx, code)
	if err != nil {
		return model.GetNextPreviousSubmissionsResponse{}, err
	}

	_, err = service.ModuleSubmissionsRepository.FindByModId(ctx, tx, course.Id, idArticle)
	if err != nil {
		return model.GetNextPreviousSubmissionsResponse{}, err
	}

	previous, err := service.ModuleSubmissionsRepository.Previous(ctx, tx, course.Id, idArticle)
	if err != nil {
		return model.GetNextPreviousSubmissionsResponse{}, err
	}

	return utils.ToModuleSubmissionsNextPreviousResponse(previous), nil
}
