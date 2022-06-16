package service

import (
	"context"
	"database/sql"
	"github.com/rg-km/final-project-engineering-12/backend/entity"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/repository"
	"github.com/rg-km/final-project-engineering-12/backend/utils"
)

type ModuleService interface {
	FindAll(ctx context.Context, limit int) ([]model.GetModuleResponse, error)
	FindAllByRelation(ctx context.Context, codeCourse string) ([]model.GetModuleResponse, error)
	FindById(ctx context.Context, codeCourse string, id int) (model.GetModuleResponse, error)
	Create(ctx context.Context, request model.CreateModuleRequest) (model.GetModuleResponse, error)
	CreateByCourse(ctx context.Context, request model.CreateModuleByCourseRequest) (model.GetModuleResponse, error)
	Update(ctx context.Context, request model.UpdateModuleRequest, codeCourse string, id int) (model.GetModuleResponse, error)
	Delete(ctx context.Context, codeCourse string, id int) error
}

type moduleService struct {
	ModuleRepository repository.ModuleRepository
	CourseRepository repository.CourseRepository
	DB               *sql.DB
}

func NewModuleService(moduleRepository *repository.ModuleRepository, courseRepository *repository.CourseRepository, db *sql.DB) ModuleService {
	return &moduleService{
		ModuleRepository: *moduleRepository,
		CourseRepository: *courseRepository,
		DB:               db,
	}
}

func (service *moduleService) FindAll(ctx context.Context, limit int) ([]model.GetModuleResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return []model.GetModuleResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	modules, err := service.ModuleRepository.FindAll(ctx, tx, limit)
	if err != nil {
		return []model.GetModuleResponse{}, err
	}

	var moduleResponses []model.GetModuleResponse
	for _, module := range modules {
		moduleResponses = append(moduleResponses, utils.ToModuleResponse(module))
	}

	return moduleResponses, nil
}

func (service *moduleService) FindAllByRelation(ctx context.Context, codeCourse string) ([]model.GetModuleResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return []model.GetModuleResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	course, err := service.CourseRepository.FindByCode(ctx, tx, codeCourse)
	if err != nil {
		return []model.GetModuleResponse{}, err
	}
	modules, err := service.ModuleRepository.FindAllByRelation(ctx, tx, course.Id)
	if err != nil {
		return []model.GetModuleResponse{}, err
	}

	var moduleResponses []model.GetModuleResponse
	for _, module := range modules {
		moduleResponses = append(moduleResponses, utils.ToModuleResponse(module))
	}

	return moduleResponses, nil
}

func (service *moduleService) FindById(ctx context.Context, codeCourse string, id int) (model.GetModuleResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetModuleResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	course, err := service.CourseRepository.FindByCode(ctx, tx, codeCourse)
	if err != nil {
		return model.GetModuleResponse{}, err
	}
	module, err := service.ModuleRepository.FindById(ctx, tx, course.Id, id)
	if err != nil {
		return model.GetModuleResponse{}, err
	}

	return utils.ToModuleResponse(module), nil
}

func (service *moduleService) Create(ctx context.Context, request model.CreateModuleRequest) (model.GetModuleResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetModuleResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	_, err = service.CourseRepository.FindById(ctx, tx, request.CourseId)
	if err != nil {
		return model.GetModuleResponse{}, err
	}

	newModule := entity.Modules{
		CourseId: request.CourseId,
		Name:     request.Name,
		IsLocked: request.IsLocked,
		Estimate: request.Estimate,
		Deadline: utils.ParseTime(request.Deadline),
	}

	module, err := service.ModuleRepository.Create(ctx, tx, newModule)
	if err != nil {
		return model.GetModuleResponse{}, err
	}

	return utils.ToModuleResponse(module), nil
}

func (service *moduleService) CreateByCourse(ctx context.Context, request model.CreateModuleByCourseRequest) (model.GetModuleResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetModuleResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	course, err := service.CourseRepository.FindByCode(ctx, tx, request.CodeCourse)
	if err != nil {
		return model.GetModuleResponse{}, err
	}

	newModule := entity.Modules{
		CourseId: course.Id,
		Name:     request.Name,
		IsLocked: request.IsLocked,
		Estimate: request.Estimate,
		Deadline: utils.ParseTime(request.Deadline),
	}

	module, err := service.ModuleRepository.Create(ctx, tx, newModule)
	if err != nil {
		return model.GetModuleResponse{}, err
	}

	return utils.ToModuleResponse(module), nil
}

func (service *moduleService) Update(ctx context.Context, request model.UpdateModuleRequest, codeCourse string, id int) (model.GetModuleResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetModuleResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	course, err := service.CourseRepository.FindByCode(ctx, tx, codeCourse)
	if err != nil {
		return model.GetModuleResponse{}, err
	}
	_, err = service.ModuleRepository.FindById(ctx, tx, course.Id, id)
	if err != nil {
		return model.GetModuleResponse{}, err
	}

	newModule := entity.Modules{
		CourseId: request.CourseId,
		Name:     request.Name,
		IsLocked: request.IsLocked,
		Estimate: request.Estimate,
		Deadline: utils.ParseTime(request.Deadline),
	}

	module, err := service.ModuleRepository.Update(ctx, tx, newModule, id)
	if err != nil {
		return model.GetModuleResponse{}, err
	}

	return utils.ToModuleResponse(module), nil
}

func (service *moduleService) Delete(ctx context.Context, codeCourse string, id int) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer utils.CommitOrRollback(tx)

	course, err := service.CourseRepository.FindByCode(ctx, tx, codeCourse)
	if err != nil {
		return err
	}
	_, err = service.ModuleRepository.FindById(ctx, tx, course.Id, id)
	if err != nil {
		return err
	}

	err = service.ModuleRepository.Delete(ctx, tx, id)
	if err != nil {
		return err
	}

	return nil
}
