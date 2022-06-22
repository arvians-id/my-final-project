package service

import (
	"context"
	"errors"
	"database/sql"
	"github.com/rg-km/final-project-engineering-12/backend/entity"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/repository"
	"github.com/rg-km/final-project-engineering-12/backend/utils"
)

type AnswerService interface {
	FindAll(ctx context.Context) ([]model.GetAnswerResponse, error)
	Create(ctx context.Context, request model.CreateAnswerRequest) (model.GetAnswerResponse, error)
	Delete(ctx context.Context, answerId int) error
	Update(ctx context.Context, request model.UpdateAnswerRequest, answerId int) (model.GetAnswerResponse, error)
	FindByUserId(ctx context.Context, userId int) ([]model.GetAnswerResponse, error)
}

type answerService struct {
	AnswerRepository		repository.AnswerRepository
	QuestionRepository 	repository.QuestionRepository
	DB         		 	    *sql.DB
}

func NewAnswerService(answerRepository *repository.AnswerRepository, questionRepository *repository.QuestionRepository, db *sql.DB) AnswerService {
	return &answerService{
		AnswerRepository: *answerRepository,
		QuestionRepository: *questionRepository,
		DB:               db,
	}
}

func (service *answerService) Create(ctx context.Context, request model.CreateAnswerRequest) (model.GetAnswerResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetAnswerResponse{}, err
	}
	defer utils.CommitOrRollback(tx)
	question, err := service.QuestionRepository.FindById(ctx, tx, request.QuestionId)
	if err != nil {
		return model.GetAnswerResponse{}, err
	}
	newAnswer := entity.Answers{
		QuestionId:  	question.Id,
		UserId:				request.UserId,
		Description: 	request.Description,
		CreatedAt:   	utils.TimeNow(),
		UpdatedAt:   	utils.TimeNow(),
	}

	answer, err := service.AnswerRepository.Create(ctx, tx, newAnswer)
	if err != nil {
		return model.GetAnswerResponse{}, err
	}

	return utils.ToAnswerResponse(answer), nil
}


func (service *answerService) FindAll(ctx context.Context) ([]model.GetAnswerResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return []model.GetAnswerResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	courses, err := service.AnswerRepository.FindAll(ctx, tx)
	if err != nil {
		return []model.GetAnswerResponse{}, err
	}

	var courseResponses []model.GetAnswerResponse
	for _, answer := range courses {
		courseResponses = append(courseResponses, utils.ToAnswerResponse(answer))
	}

	return courseResponses, nil
}


func (service *answerService) Delete(ctx context.Context, answerId int) error {
	// userId := 11; // ini nanti akan diubah pakai data auth user-id dari middleware auth
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer utils.CommitOrRollback(tx)
	
	getAnswer, err := service.AnswerRepository.FindById(ctx, tx, answerId)
	if err != nil {
		return err
	}

	// if getAnswer.UserId != userId {
	// 	return errors.New("access not allowed")
	// }
	
	err = service.AnswerRepository.Delete(ctx, tx, getAnswer.Id)
	if err != nil {
		return err
	}

	return nil
}

func (service *answerService) Update(ctx context.Context, request model.UpdateAnswerRequest, answerId int) (model.GetAnswerResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetAnswerResponse{}, err
	}
	defer utils.CommitOrRollback(tx)
	getAnswer, err := service.AnswerRepository.FindById(ctx, tx, answerId)
	if err != nil {
		return model.GetAnswerResponse{}, err
	}

	if getAnswer.UserId != request.UserId {
		return model.GetAnswerResponse{}, errors.New("access not allowed")
	}

	question, err := service.QuestionRepository.FindById(ctx, tx, request.QuestionId)
	if err != nil {
		return model.GetAnswerResponse{}, err
	}

	newAnswer := entity.Answers{
		QuestionId:  	question.Id,
		UserId:				getAnswer.UserId,
		Description: 	request.Description,
		CreatedAt:  	getAnswer.CreatedAt,
		UpdatedAt:   	utils.TimeNow(),
	}

	_, err = service.AnswerRepository.Update(ctx, tx, newAnswer, answerId)
	if err != nil {
		return model.GetAnswerResponse{}, err
	}

	getAnswerUpdate, err := service.AnswerRepository.FindById(ctx, tx, answerId)
	if err != nil {
		return model.GetAnswerResponse{}, err
	}

	return utils.ToAnswerResponse(getAnswerUpdate), nil
}


func (service *answerService) FindByUserId(ctx context.Context, userId int) ([]model.GetAnswerResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return []model.GetAnswerResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	courses, err := service.AnswerRepository.FindByUserId(ctx, tx, userId)
	if err != nil {
		return []model.GetAnswerResponse{}, err
	}

	var courseResponses []model.GetAnswerResponse
	for _, answer := range courses {
		courseResponses = append(courseResponses, utils.ToAnswerResponse(answer))
	}

	return courseResponses, nil
}