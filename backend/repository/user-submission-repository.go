package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/rg-km/final-project-engineering-12/backend/entity"
)

type UserSubmissionsRepository interface {
	OnlyCreate(ctx context.Context, tx *sql.Tx, userId int, moduleSubmissionId int) error
	UpdateFile(ctx context.Context, tx *sql.Tx, userSubmission entity.UserSubmissions) (entity.UserSubmissions, error)
	UpdateGrade(ctx context.Context, tx *sql.Tx, userSubmission entity.UserSubmissions) error
	FindUserSubmissionByOther(ctx context.Context, tx *sql.Tx, userSubmission entity.UserSubmissions) (entity.UserSubmissions, error)
	FindUserSubmissionById(ctx context.Context, tx *sql.Tx, id int) (entity.UserSubmissions, error)
}

type userSubmissionsRepository struct {
}

func NewUserSubmissionsRepository() UserSubmissionsRepository {
	return &userSubmissionsRepository{}
}

func (repository *userSubmissionsRepository) OnlyCreate(ctx context.Context, tx *sql.Tx, userId int, moduleSubmissionId int) error {
	query := `INSERT INTO user_submissions(user_id, module_submission_id) VALUES(?,?)`
	_, err := tx.ExecContext(
		ctx,
		query,
		userId,
		moduleSubmissionId,
	)
	if err != nil {
		return err
	}

	return nil
}

func (repository *userSubmissionsRepository) UpdateFile(ctx context.Context, tx *sql.Tx, userSubmission entity.UserSubmissions) (entity.UserSubmissions, error) {
	query := `UPDATE user_submissions SET file = ? WHERE user_id = ? AND module_submission_id = ?`
	_, err := tx.ExecContext(
		ctx,
		query,
		userSubmission.File,
		userSubmission.UserId,
		userSubmission.ModuleSubmissionId,
	)
	if err != nil {
		return entity.UserSubmissions{}, err
	}

	return userSubmission, nil
}

func (repository *userSubmissionsRepository) UpdateGrade(ctx context.Context, tx *sql.Tx, userSubmission entity.UserSubmissions) error {
	query := `UPDATE user_submissions SET grade = ? WHERE id = ?`
	_, err := tx.ExecContext(
		ctx,
		query,
		userSubmission.Grade,
		userSubmission.Id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (repository *userSubmissionsRepository) FindUserSubmissionByOther(ctx context.Context, tx *sql.Tx, userSubmission entity.UserSubmissions) (entity.UserSubmissions, error) {
	query := `SELECT * FROM user_submissions WHERE user_id = ? AND module_submission_id = ?`
	queryContext, err := tx.QueryContext(ctx, query, userSubmission.UserId, userSubmission.ModuleSubmissionId)
	if err != nil {
		return entity.UserSubmissions{}, err
	}
	defer func(queryContext *sql.Rows) {
		err := queryContext.Close()
		if err != nil {
			return
		}
	}(queryContext)

	var modsub entity.UserSubmissions
	if queryContext.Next() {
		err := queryContext.Scan(
			&modsub.Id,
			&modsub.UserId,
			&modsub.ModuleSubmissionId,
			&modsub.File,
			&modsub.Grade,
		)
		if err != nil {
			return entity.UserSubmissions{}, err
		}

		return modsub, nil
	}

	return modsub, errors.New("user submission not found")
}

func (repository *userSubmissionsRepository) FindUserSubmissionById(ctx context.Context, tx *sql.Tx, id int) (entity.UserSubmissions, error) {
	query := `SELECT * FROM user_submissions WHERE id = ?`
	queryContext, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return entity.UserSubmissions{}, err
	}
	defer func(queryContext *sql.Rows) {
		err := queryContext.Close()
		if err != nil {
			return
		}
	}(queryContext)

	var modsub entity.UserSubmissions
	if queryContext.Next() {
		err := queryContext.Scan(
			&modsub.Id,
			&modsub.UserId,
			&modsub.ModuleSubmissionId,
			&modsub.File,
			&modsub.Grade,
		)
		if err != nil {
			return entity.UserSubmissions{}, err
		}

		return modsub, nil
	}

	return modsub, errors.New("user submission not found")
}
