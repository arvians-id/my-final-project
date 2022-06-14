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
	FindAll(ctx context.Context) ([]model.GetModuleSubmissionsResponse, error)
	FindByModId(ctx context.Context, code string) (model.GetModuleSubmissionsResponse, error)
	Create(ctx context.Context, request model.CreateModuleSubmissionsRequest) (model.GetModuleSubmissionsResponse, error)
	Update(ctx context.Context, request model.UpdateModuleSubmissionsRequest, code string) (model.GetModuleSubmissionsResponse, error)
	Delete(ctx context.Context, code string) error
}

type moduleSubmissionsService struct {
	ModuleSubmissionsRepository repository.ModuleSubmissionsRepository
	DB                          *sql.DB
}

func NewModuleSubmissionsService(moduleSubmissionsRepository *repository.ModuleSubmissionsRepository, db *sql.DB) ModuleSubmissionsService {
	return &moduleSubmissionsService{
		ModuleSubmissionsRepository: *moduleSubmissionsRepository,
		DB:                          db,
	}
}

func (service *moduleSubmissionsService) FindAll(ctx context.Context) ([]model.GetModuleSubmissionsResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return []model.GetModuleSubmissionsResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	modsubs, err := service.ModuleSubmissionsRepository.FindAll(ctx, tx)
	if err != nil {
		return []model.GetModuleSubmissionsResponse{}, err
	}

	var modsubResponses []model.GetModuleSubmissionsResponse
	for _, modsub := range modsubs {
		modsubResponses = append(modsubResponses, utils.ToModuleSubmissionsResponse(modsub))
	}

	return modsubResponses, nil
}

func (service *moduleSubmissionsService) FindByModId(ctx context.Context, code string) (model.GetModuleSubmissionsResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetModuleSubmissionsResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	modsub, err := service.ModuleSubmissionsRepository.FindByModId(ctx, tx, code)
	if err != nil {
		return model.GetModuleSubmissionsResponse{}, err
	}

	return utils.ToModuleSubmissionsResponse(modsub), nil
}

func (service *moduleSubmissionsService) Create(ctx context.Context, request model.CreateModuleSubmissionsRequest) (model.GetModuleSubmissionsResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetModuleSubmissionsResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	newModsub := entity.ModuleSubmissions{
		Id:       request.Id,
		ModuleId: request.ModuleId,
		File:     request.File,
		Type:     request.Type,
		MaxSize:  request.MaxSize,
	}

	modsub, err := service.ModuleSubmissionsRepository.Create(ctx, tx, newModsub)
	if err != nil {
		return model.GetModuleSubmissionsResponse{}, err
	}

	return utils.ToModuleSubmissionsResponse(modsub), nil
}

func (service *moduleSubmissionsService) Update(ctx context.Context, request model.UpdateModuleSubmissionsRequest, code string) (model.GetModuleSubmissionsResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetModuleSubmissionsResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	getModsub, err := service.ModuleSubmissionsRepository.FindByModId(ctx, tx, code)
	if err != nil {
		return model.GetModuleSubmissionsResponse{}, err
	}

	newModsub := entity.ModuleSubmissions{
		Id:       getModsub.Id,
		ModuleId: request.ModuleId,
		File:     request.File,
		Type:     request.Type,
		MaxSize:  request.MaxSize,
	}

	modsub, err := service.ModuleSubmissionsRepository.Update(ctx, tx, newModsub)
	if err != nil {
		return model.GetModuleSubmissionsResponse{}, err
	}

	return utils.ToModuleSubmissionsResponse(modsub), nil
}

func (service *moduleSubmissionsService) Delete(ctx context.Context, code string) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer utils.CommitOrRollback(tx)

	getModsub, err := service.ModuleSubmissionsRepository.FindByModId(ctx, tx, code)
	if err != nil {
		return err
	}

	err = service.ModuleSubmissionsRepository.Delete(ctx, tx, utils.ToString(getModsub.ModuleId))
	if err != nil {
		return err
	}

	return nil
}
