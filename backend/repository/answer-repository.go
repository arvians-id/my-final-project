package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/rg-km/final-project-engineering-12/backend/entity"
)

type AnswerRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) ([]entity.Answers, error)
	Create(ctx context.Context, tx *sql.Tx, answer entity.Answers) (entity.Answers, error)
	Delete(ctx context.Context, tx *sql.Tx, answerId int) error
	FindById(ctx context.Context, tx *sql.Tx, answerId int) (entity.Answers, error)
	Update(ctx context.Context, tx *sql.Tx, answer entity.Answers, answerId int) (entity.Answers, error)
	FindByUserId(ctx context.Context, tx *sql.Tx, userId int) ([]entity.Answers, error)
}

type answerRepository struct {
}

func NewAnswerRepository() AnswerRepository {
	return &answerRepository{}
}

func (repository *answerRepository) Create(ctx context.Context, tx *sql.Tx, answer entity.Answers) (entity.Answers, error) {
	query := `INSERT INTO answers(question_id, user_id, description, created_at, updated_at) VALUES(?,?,?,?,?)`

	queryContext, err := tx.ExecContext(
		ctx,
		query,
		answer.QuestionId,
		answer.UserId,
		answer.Description,
		answer.CreatedAt,
		answer.UpdatedAt,
	)
	if err != nil {
		return entity.Answers{}, err
	}

	id, err := queryContext.LastInsertId()
	if err != nil {
		return entity.Answers{}, err
	}
	answer.Id = int(id)

	return answer, nil
}

func (repository *answerRepository) FindAll(ctx context.Context, tx *sql.Tx) ([]entity.Answers, error) {
	query := `SELECT * FROM answers ORDER BY created_at DESC`
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

	var answers []entity.Answers
	for queryContext.Next() {
		var answer entity.Answers
		err := queryContext.Scan(
			&answer.Id,
			&answer.QuestionId,
			&answer.UserId,
			&answer.Description,
			&answer.CreatedAt,
			&answer.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		answers = append(answers, answer)
	}

	return answers, nil
}

func (repository *answerRepository) Delete(ctx context.Context, tx *sql.Tx, answerId int) error {
	query := "DELETE FROM answers WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, answerId)
	if err != nil {
		return err
	}

	return nil
}

func (repository *answerRepository) FindById(ctx context.Context, tx *sql.Tx,  answerId int) (entity.Answers, error) {
	query := `SELECT * FROM answers WHERE id = ?`
	queryContext, err := tx.QueryContext(ctx, query, answerId)
	if err != nil {
		return entity.Answers{}, err
	}
	defer func(queryContext *sql.Rows) {
		err := queryContext.Close()
		if err != nil {
			return
		}
	}(queryContext)

	var answer entity.Answers
	if queryContext.Next() {
		err := queryContext.Scan(
			&answer.Id,
			&answer.QuestionId,
			&answer.UserId,
			&answer.Description,
			&answer.CreatedAt,
			&answer.UpdatedAt,
		)
		if err != nil {
			return answer, err
		}

		return answer, nil
	}

	return answer, errors.New("answer not found")
}

func (repository *answerRepository) Update(ctx context.Context, tx *sql.Tx, answer entity.Answers, answerId int) (entity.Answers, error) {

	query := `UPDATE answers SET question_id = ?, description = ?, updated_at = ? WHERE id = ?`
	_, err := tx.ExecContext(
		ctx,
		query,
		answer.QuestionId,
		answer.Description,
		answer.UpdatedAt,
		answerId,
	)
	if err != nil {
		return entity.Answers{}, err
	}

	return answer, nil
}

func (repository *answerRepository) FindByUserId(ctx context.Context, tx *sql.Tx, userId int) ([]entity.Answers, error) {
	query := `SELECT * FROM answers WHERE user_id = ? ORDER BY created_at DESC`
	queryContext, err := tx.QueryContext(ctx, query, userId)
	if err != nil {
		return []entity.Answers{}, err
	}
	defer func(queryContext *sql.Rows) {
		err := queryContext.Close()
		if err != nil {
			return
		}
	}(queryContext)

	var answers []entity.Answers
	for queryContext.Next() {
		var answer entity.Answers
		err := queryContext.Scan(
			&answer.Id,
			&answer.QuestionId,
			&answer.UserId,
			&answer.Description,
			&answer.CreatedAt,
			&answer.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		answers = append(answers, answer)
	}

	return answers, nil
}