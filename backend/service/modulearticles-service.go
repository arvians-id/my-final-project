package service

import (
	"context"
	"database/sql"

	"github.com/rg-km/final-project-engineering-12/backend/entity"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/repository"
	"github.com/rg-km/final-project-engineering-12/backend/utils"
)

type ModuleArticlesService interface {
	FindAll(ctx context.Context) ([]model.GetModuleArticlesResponse, error)
	FindByModId(ctx context.Context, code string) (model.GetModuleArticlesResponse, error)
	Create(ctx context.Context, request model.CreateModuleArticlesRequest) (model.GetModuleArticlesResponse, error)
	Update(ctx context.Context, request model.UpdateModuleArticlesRequest, code string) (model.GetModuleArticlesResponse, error)
	Delete(ctx context.Context, code string) error
}

type moduleArticlesService struct {
	ModuleArticlesRepository repository.ModuleArticlesRepository
	DB                       *sql.DB
}

func NewModuleArticlesService(moduleArticlesRepository *repository.ModuleArticlesRepository, db *sql.DB) ModuleArticlesService {
	return &moduleArticlesService{
		ModuleArticlesRepository: *moduleArticlesRepository,
		DB:                       db,
	}
}

func (service *moduleArticlesService) FindAll(ctx context.Context) ([]model.GetModuleArticlesResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return []model.GetModuleArticlesResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	ModArs, err := service.ModuleArticlesRepository.FindAll(ctx, tx)
	if err != nil {
		return []model.GetModuleArticlesResponse{}, err
	}

	var ModArResponses []model.GetModuleArticlesResponse
	for _, ModAr := range ModArs {
		ModArResponses = append(ModArResponses, utils.ToModuleArticlesResponse(ModAr))
	}

	return ModArResponses, nil
}

func (service *moduleArticlesService) FindByModId(ctx context.Context, code string) (model.GetModuleArticlesResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetModuleArticlesResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	ModAr, err := service.ModuleArticlesRepository.FindByModId(ctx, tx, code)
	if err != nil {
		return model.GetModuleArticlesResponse{}, err
	}

	return utils.ToModuleArticlesResponse(ModAr), nil
}

func (service *moduleArticlesService) Create(ctx context.Context, request model.CreateModuleArticlesRequest) (model.GetModuleArticlesResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetModuleArticlesResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	newModArs := entity.ModuleArticles{
		ModuleId: request.ModuleId,
		Content:  request.Content,
	}

	ModAr, err := service.ModuleArticlesRepository.Create(ctx, tx, newModArs)
	if err != nil {
		return model.GetModuleArticlesResponse{}, err
	}

	return utils.ToModuleArticlesResponse(ModAr), nil
}

func (service *moduleArticlesService) Update(ctx context.Context, request model.UpdateModuleArticlesRequest, code string) (model.GetModuleArticlesResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetModuleArticlesResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	getModArs, err := service.ModuleArticlesRepository.FindByModId(ctx, tx, code)
	if err != nil {
		return model.GetModuleArticlesResponse{}, err
	}

	newModArs := entity.ModuleArticles{
		ModuleId: getModArs.ModuleId,
		Content:  request.Content,
	}

	ModAr, err := service.ModuleArticlesRepository.Update(ctx, tx, newModArs)
	if err != nil {
		return model.GetModuleArticlesResponse{}, err
	}

	return utils.ToModuleArticlesResponse(ModAr), nil
}

func (service *moduleArticlesService) Delete(ctx context.Context, code string) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer utils.CommitOrRollback(tx)

	getModArs, err := service.ModuleArticlesRepository.FindByModId(ctx, tx, code)
	if err != nil {
		return err
	}

	err = service.ModuleArticlesRepository.Delete(ctx, tx, utils.ToString(getModArs.ModuleId))
	if err != nil {
		return err
	}

	return nil
}
