package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/rg-km/final-project-engineering-12/backend/entity"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/repository"
	"github.com/rg-km/final-project-engineering-12/backend/utils"
)

type QuestionService interface {
	FindAll(ctx context.Context) ([]model.GetQuestionRelationResponse, error)
	Create(ctx context.Context, request model.CreateQuestionRequest) (model.GetQuestionResponse, error)
	Delete(ctx context.Context, questionId int) error
	Update(ctx context.Context, request model.UpdateQuestionRequest, questionId int) (model.GetQuestionRelationResponse, error)
	FindByUserId(ctx context.Context, userId int) ([]model.GetQuestionRelationResponse, error)
	FindById(ctx context.Context, questionId int) (model.GetQuestionRelationResponse, error)
}

type questionService struct {
	QuestionRepository repository.QuestionRepository
	UserRepository    repository.UserRepository
	DB                 *sql.DB
}

func NewQuestionService(questionRepository *repository.QuestionRepository, userRepository *repository.UserRepository, db *sql.DB) QuestionService {
	return &questionService{
		QuestionRepository: *questionRepository,
		UserRepository:    *userRepository,
		DB:                 db,
	}
}

func (service *questionService) Create(ctx context.Context, request model.CreateQuestionRequest) (model.GetQuestionResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetQuestionResponse{}, err
	}
	defer utils.CommitOrRollback(tx)
	user, err := service.UserRepository.GetUserByID(ctx, tx, request.UserId)

	if err != nil {
		return model.GetQuestionResponse{}, err
	}
	
	if(user.Id == 0 ){
		return model.GetQuestionResponse{}, errors.New("user not found")
	}

	newQuestion := entity.Questions{
		UserId:      request.UserId,
		CourseId:    request.CourseId,
		Title:       request.Title,
		Tags:        request.Tags,
		Description: request.Description,
		CreatedAt:   utils.TimeNow(),
		UpdatedAt:   utils.TimeNow(),
	}

	question, err := service.QuestionRepository.Create(ctx, tx, newQuestion)
	if err != nil {
		return model.GetQuestionResponse{}, err
	}

	return utils.ToQuestionResponse(question), nil
}

func (service *questionService) FindAll(ctx context.Context) ([]model.GetQuestionRelationResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return []model.GetQuestionRelationResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	courses, err := service.QuestionRepository.FindAll(ctx, tx)
	if err != nil {
		return []model.GetQuestionRelationResponse{}, err
	}

	var courseResponses []model.GetQuestionRelationResponse
	for _, question := range courses {
		courseResponses = append(courseResponses, utils.ToQuestionRelationResponse(question))
	}

	return courseResponses, nil
}

func (service *questionService) Delete(ctx context.Context, questionId int) error {
	// userId := 11; // ini nanti akan diubah pakai data auth user-id dari middleware auth
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer utils.CommitOrRollback(tx)

	getQuestions, err := service.QuestionRepository.FindById(ctx, tx, questionId)
	if err != nil {
		return err
	}

	// if getQuestions.UserId != userId {
	// 	return errors.New("access not allowed")
	// }

	err = service.QuestionRepository.Delete(ctx, tx, getQuestions.Id)
	if err != nil {
		return err
	}

	return nil
}

func (service *questionService) Update(ctx context.Context, request model.UpdateQuestionRequest, questionId int) (model.GetQuestionRelationResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetQuestionRelationResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	getQuestions, err := service.QuestionRepository.FindById(ctx, tx, questionId)
	if err != nil {
		return model.GetQuestionRelationResponse{}, err
	}

	if getQuestions.UserId != request.UserId {
		return model.GetQuestionRelationResponse{}, errors.New("access not allowed")
	}

	newQuestion := entity.Questions{
		UserId:      getQuestions.UserId,
		CourseId:    request.CourseId,
		Title:       request.Title,
		Tags:        request.Tags,
		Description: request.Description,
		CreatedAt:   getQuestions.CreatedAt,
		UpdatedAt:   utils.TimeNow(),
	}

	_, err = service.QuestionRepository.Update(ctx, tx, newQuestion, questionId)
	if err != nil {
		return model.GetQuestionRelationResponse{}, err
	}

	getQuestionsUpdate, err := service.QuestionRepository.FindById(ctx, tx, questionId)
	if err != nil {
		return model.GetQuestionRelationResponse{}, err
	}

	return utils.ToQuestionRelationResponse(getQuestionsUpdate), nil
}

func (service *questionService) FindByUserId(ctx context.Context, userId int) ([]model.GetQuestionRelationResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return []model.GetQuestionRelationResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	courses, err := service.QuestionRepository.FindByUserId(ctx, tx, userId)
	if err != nil {
		return []model.GetQuestionRelationResponse{}, err
	}

	var courseResponses []model.GetQuestionRelationResponse
	for _, question := range courses {
		courseResponses = append(courseResponses, utils.ToQuestionRelationResponse(question))
	}

	return courseResponses, nil
}

func (service *questionService) FindById(ctx context.Context, questionId int) (model.GetQuestionRelationResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return model.GetQuestionRelationResponse{}, err
	}
	defer utils.CommitOrRollback(tx)

	questions, err := service.QuestionRepository.FindById(ctx, tx, questionId)
	if err != nil {
		return model.GetQuestionRelationResponse{}, err
	}

	return utils.ToQuestionRelationResponse(questions), nil
}

