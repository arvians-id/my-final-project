package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/rg-km/final-project-engineering-12/backend/entity"
)

type QuestionRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) ([]entity.Questions, error)
	Create(ctx context.Context, tx *sql.Tx, question entity.Questions) (entity.Questions, error)
	Delete(ctx context.Context, tx *sql.Tx, questionId int) error
	FindById(ctx context.Context, tx *sql.Tx, questionId int) (entity.Questions, error)
	Update(ctx context.Context, tx *sql.Tx, question entity.Questions, questionId int) (entity.Questions, error)
	FindByUserId(ctx context.Context, tx *sql.Tx, userId int) ([]entity.Questions, error)
}

type questionRepository struct {
}

func NewQuestionRepository() QuestionRepository {
	return &questionRepository{}
}

func (repository *questionRepository) Create(ctx context.Context, tx *sql.Tx, question entity.Questions) (entity.Questions, error) {
	query := `INSERT INTO questions(user_id, module_id, title, tags, description, created_at, updated_at) VALUES(?,?,?,?,?,?,?)`

	queryContext, err := tx.ExecContext(
		ctx,
		query,
		question.UserId,
		question.ModuleId,
		question.Title,
		question.Tags,
		question.Description,
		question.CreatedAt,
		question.UpdatedAt,
	)
	if err != nil {
		return entity.Questions{}, err
	}

	id, err := queryContext.LastInsertId()
	if err != nil {
		return entity.Questions{}, err
	}
	question.Id = int(id)

	return question, nil
}

func (repository *questionRepository) FindAll(ctx context.Context, tx *sql.Tx) ([]entity.Questions, error) {
	query := `SELECT * FROM questions ORDER BY created_at DESC`
	queryContext, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func(queryContext *sql.Rows) {
		err := queryContext.Close()
		if err != nil {
			return
		}
	}(queryContext)

	var questions []entity.Questions
	for queryContext.Next() {
		var question entity.Questions
		err := queryContext.Scan(
			&question.Id,
			&question.ModuleId,
			&question.UserId,
			&question.Title,
			&question.Tags,
			&question.Description,
			&question.CreatedAt,
			&question.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		questions = append(questions, question)
	}

	return questions, nil
}

func (repository *questionRepository) Delete(ctx context.Context, tx *sql.Tx, questionId int) error {
	query := "DELETE FROM questions WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, questionId)
	if err != nil {
		return err
	}

	return nil
}

func (repository *questionRepository) FindById(ctx context.Context, tx *sql.Tx,  questionId int) (entity.Questions, error) {
	query := `SELECT * FROM questions WHERE id = ?`
	queryContext, err := tx.QueryContext(ctx, query, questionId)
	if err != nil {
		return entity.Questions{}, err
	}
	defer func(queryContext *sql.Rows) {
		err := queryContext.Close()
		if err != nil {
			return
		}
	}(queryContext)

	var question entity.Questions
	if queryContext.Next() {
		err := queryContext.Scan(
			&question.Id,
			&question.ModuleId,
			&question.UserId,
			&question.Title,
			&question.Tags,
			&question.Description,
			&question.CreatedAt,
			&question.UpdatedAt,
		)
		if err != nil {
			return question, err
		}

		return question, nil
	}

	return question, errors.New("question not found")
}

func (repository *questionRepository) Update(ctx context.Context, tx *sql.Tx, question entity.Questions, questionId int) (entity.Questions, error) {

	query := `UPDATE questions SET module_id = ?, title = ?, tags = ?, description = ?, updated_at = ? WHERE id = ?`
	_, err := tx.ExecContext(
		ctx,
		query,
		question.ModuleId,
		question.Title,
		question.Tags,
		question.Description,
		question.UpdatedAt,
		questionId,
	)
	if err != nil {
		return entity.Questions{}, err
	}

	return question, nil
}

func (repository *questionRepository) FindByUserId(ctx context.Context, tx *sql.Tx, userId int) ([]entity.Questions, error) {
	query := `SELECT * FROM questions WHERE user_id = ? ORDER BY created_at DESC`
	queryContext, err := tx.QueryContext(ctx, query, userId)
	if err != nil {
		return []entity.Questions{}, err
	}
	defer func(queryContext *sql.Rows) {
		err := queryContext.Close()
		if err != nil {
			return
		}
	}(queryContext)

	var questions []entity.Questions
	for queryContext.Next() {
		var question entity.Questions
		err := queryContext.Scan(
			&question.Id,
			&question.ModuleId,
			&question.UserId,
			&question.Title,
			&question.Tags,
			&question.Description,
			&question.CreatedAt,
			&question.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		questions = append(questions, question)
	}

	return questions, nil
}