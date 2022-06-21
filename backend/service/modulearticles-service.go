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
	FindAll(ctx context.Context, code string) ([]model.GetModuleArticlesResponse, error)
	FindByModId(ctx context.Context, code string, idArticle int) (model.GetModuleArticlesResponse, error)
	Create(ctx context.Context, request model.CreateModuleArticlesRequest, code string) (model.GetModuleArticlesResponse, error)
	Update(ctx context.Context, request model.UpdateModuleArticlesRequest, code string, idArticle int) (model.GetModuleArticlesResponse, error)
	Delete(ctx context.Context, code string, idArticle int) error
	Next(ctx context.Context, code string, idArticle int) (model.GetNextPreviousArticlesResponse, error)
	Previous(ctx context.Context, code string, idArticle int) (model.GetNextPreviousArticlesResponse, error)
}

type moduleArticlesService struct {
	ModuleArticlesRepository repository.ModuleArticlesRepository
	CourseRepository         repository.CourseRepository
	DB                       *sql.DB
}

func NewModuleArticlesService(moduleArticlesRepository *repository.ModuleArticlesRepository, courseRepository *repository.CourseRepository, db *sql.DB) ModuleArticlesService {
	return &moduleArticlesService{
		ModuleArticlesRepository: *moduleArticlesRepository,
		CourseRepository:         *courseRepository,
		DB:                       db,
	}
}

func (service *moduleArticlesService) FindAll(ctx context.Context, code string) ([]model.GetModuleArticlesResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return []model.GetModuleArticlesResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	course, err := service.CourseRepository.FindByCode(ctx, tx, code)
	if err != nil {
		return []model.GetModuleArticlesResponse{}, err
	}

	ModArs, err := service.ModuleArticlesRepository.FindAll(ctx, tx, course.Id)
	if err != nil {
		return []model.GetModuleArticlesResponse{}, err
	}

	var ModArResponses []model.GetModuleArticlesResponse
	for _, ModAr := range ModArs {
		ModArResponses = append(ModArResponses, utils.ToModuleArticlesResponse(ModAr))
	}

	return ModArResponses, nil
}

func (service *moduleArticlesService) FindByModId(ctx context.Context, code string, idArticle int) (model.GetModuleArticlesResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetModuleArticlesResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	course, err := service.CourseRepository.FindByCode(ctx, tx, code)
	if err != nil {
		return model.GetModuleArticlesResponse{}, err
	}

	ModAr, err := service.ModuleArticlesRepository.FindByModId(ctx, tx, course.Id, idArticle)
	if err != nil {
		return model.GetModuleArticlesResponse{}, err
	}

	return utils.ToModuleArticlesResponse(ModAr), nil
}

func (service *moduleArticlesService) Create(ctx context.Context, request model.CreateModuleArticlesRequest, code string) (model.GetModuleArticlesResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetModuleArticlesResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	course, err := service.CourseRepository.FindByCode(ctx, tx, code)
	if err != nil {
		return model.GetModuleArticlesResponse{}, err
	}

	newModArs := entity.ModuleArticles{
		CourseId: course.Id,
		Name:     request.Name,
		Content:  request.Content,
		Estimate: request.Estimate,
	}

	ModAr, err := service.ModuleArticlesRepository.Create(ctx, tx, newModArs)
	if err != nil {
		return model.GetModuleArticlesResponse{}, err
	}

	return utils.ToModuleArticlesResponse(ModAr), nil
}

func (service *moduleArticlesService) Update(ctx context.Context, request model.UpdateModuleArticlesRequest, code string, idArticle int) (model.GetModuleArticlesResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetModuleArticlesResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	course, err := service.CourseRepository.FindByCode(ctx, tx, code)
	if err != nil {
		return model.GetModuleArticlesResponse{}, err
	}

	_, err = service.ModuleArticlesRepository.FindByModId(ctx, tx, course.Id, idArticle)
	if err != nil {
		return model.GetModuleArticlesResponse{}, err
	}

	newModArs := entity.ModuleArticles{
		CourseId: course.Id,
		Name:     request.Name,
		Content:  request.Content,
		Estimate: request.Estimate,
	}

	ModAr, err := service.ModuleArticlesRepository.Update(ctx, tx, newModArs, idArticle)
	if err != nil {
		return model.GetModuleArticlesResponse{}, err
	}

	return utils.ToModuleArticlesResponse(ModAr), nil
}

func (service *moduleArticlesService) Delete(ctx context.Context, code string, idArticle int) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer utils.CommitOrRollback(tx)

	course, err := service.CourseRepository.FindByCode(ctx, tx, code)
	if err != nil {
		return err
	}

	_, err = service.ModuleArticlesRepository.FindByModId(ctx, tx, course.Id, idArticle)
	if err != nil {
		return err
	}

	err = service.ModuleArticlesRepository.Delete(ctx, tx, idArticle)
	if err != nil {
		return err
	}

	return nil
}

func (service *moduleArticlesService) Next(ctx context.Context, code string, idArticle int) (model.GetNextPreviousArticlesResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetNextPreviousArticlesResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	course, err := service.CourseRepository.FindByCode(ctx, tx, code)
	if err != nil {
		return model.GetNextPreviousArticlesResponse{}, err
	}

	_, err = service.ModuleArticlesRepository.FindByModId(ctx, tx, course.Id, idArticle)
	if err != nil {
		return model.GetNextPreviousArticlesResponse{}, err
	}

	next, err := service.ModuleArticlesRepository.Next(ctx, tx, course.Id, idArticle)
	if err != nil {
		return model.GetNextPreviousArticlesResponse{}, err
	}

	return utils.ToModuleArticlesNextPreviousResponse(next), nil
}

func (service *moduleArticlesService) Previous(ctx context.Context, code string, idArticle int) (model.GetNextPreviousArticlesResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetNextPreviousArticlesResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	course, err := service.CourseRepository.FindByCode(ctx, tx, code)
	if err != nil {
		return model.GetNextPreviousArticlesResponse{}, err
	}

	_, err = service.ModuleArticlesRepository.FindByModId(ctx, tx, course.Id, idArticle)
	if err != nil {
		return model.GetNextPreviousArticlesResponse{}, err
	}

	previous, err := service.ModuleArticlesRepository.Previous(ctx, tx, course.Id, idArticle)
	if err != nil {
		return model.GetNextPreviousArticlesResponse{}, err
	}

	return utils.ToModuleArticlesNextPreviousResponse(previous), nil
}
