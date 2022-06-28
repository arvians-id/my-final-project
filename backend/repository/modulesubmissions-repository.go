package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/rg-km/final-project-engineering-12/backend/entity"
)

type ModuleSubmissionsRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx, idCourse int) ([]entity.ModuleSubmissions, error)
	FindByModId(ctx context.Context, tx *sql.Tx, idCourse int, idSubmission int) (entity.ModuleSubmissions, error)
	Create(ctx context.Context, tx *sql.Tx, modsub entity.ModuleSubmissions) (entity.ModuleSubmissions, error)
	Update(ctx context.Context, tx *sql.Tx, modsub entity.ModuleSubmissions, idSubmission int) (entity.ModuleSubmissions, error)
	Delete(ctx context.Context, tx *sql.Tx, idSubmission int) error
	Next(ctx context.Context, tx *sql.Tx, idCourse int, idSubmission int) (entity.NextPreviousModuleSubmissions, error)
	Previous(ctx context.Context, tx *sql.Tx, idCourse int, idSubmission int) (entity.NextPreviousModuleSubmissions, error)
}

type moduleSubmissionsRepository struct {
}

func NewModuleSubmissionsRepository() ModuleSubmissionsRepository {
	return &moduleSubmissionsRepository{}
}

func (repository *moduleSubmissionsRepository) FindAll(ctx context.Context, tx *sql.Tx, idCourse int) ([]entity.ModuleSubmissions, error) {
	query := `SELECT * FROM module_submissions WHERE course_id = ?`
	queryContext, err := tx.QueryContext(ctx, query, idCourse)
	if err != nil {
		return nil, err
	}
	defer func(queryContext *sql.Rows) {
		err := queryContext.Close()
		if err != nil {
			return
		}
	}(queryContext)

	var modsubs []entity.ModuleSubmissions
	for queryContext.Next() {
		var modsub entity.ModuleSubmissions
		err := queryContext.Scan(
			&modsub.Id,
			&modsub.CourseId,
			&modsub.Name,
			&modsub.Description,
			&modsub.Deadline,
		)
		if err != nil {
			return nil, err
		}

		modsubs = append(modsubs, modsub)
	}

	return modsubs, nil
}

func (repository *moduleSubmissionsRepository) FindByModId(ctx context.Context, tx *sql.Tx, idCourse int, idSubmission int) (entity.ModuleSubmissions, error) {
	query := `SELECT * FROM module_submissions WHERE course_id = ? AND id = ?`
	queryContext, err := tx.QueryContext(ctx, query, idCourse, idSubmission)
	if err != nil {
		return entity.ModuleSubmissions{}, err
	}
	defer func(queryContext *sql.Rows) {
		err := queryContext.Close()
		if err != nil {
			return
		}
	}(queryContext)

	var modsub entity.ModuleSubmissions
	if queryContext.Next() {
		err := queryContext.Scan(
			&modsub.Id,
			&modsub.CourseId,
			&modsub.Name,
			&modsub.Description,
			&modsub.Deadline,
		)
		if err != nil {
			return entity.ModuleSubmissions{}, err
		}

		return modsub, nil
	}

	return modsub, errors.New("submission not found")
}

func (repository *moduleSubmissionsRepository) Create(ctx context.Context, tx *sql.Tx, modsub entity.ModuleSubmissions) (entity.ModuleSubmissions, error) {
	query := `INSERT INTO module_submissions(course_id, name, description, deadline) VALUES(?,?,?,?)`
	queryContext, err := tx.ExecContext(
		ctx,
		query,
		modsub.CourseId,
		modsub.Name,
		modsub.Description,
		modsub.Deadline,
	)
	if err != nil {
		return entity.ModuleSubmissions{}, err
	}

	id, err := queryContext.LastInsertId()
	if err != nil {
		return entity.ModuleSubmissions{}, err
	}
	modsub.Id = int(id)

	return modsub, nil
}

func (repository *moduleSubmissionsRepository) Update(ctx context.Context, tx *sql.Tx, modsub entity.ModuleSubmissions, idSubmission int) (entity.ModuleSubmissions, error) {
	query := `UPDATE module_submissions SET name = ?, description = ?, deadline = ? WHERE id = ?`
	_, err := tx.ExecContext(
		ctx,
		query,
		modsub.Name,
		modsub.Description,
		modsub.Deadline,
		idSubmission,
	)
	if err != nil {
		return entity.ModuleSubmissions{}, err
	}

	return modsub, nil
}

func (repository *moduleSubmissionsRepository) Delete(ctx context.Context, tx *sql.Tx, idSubmission int) error {
	query := "DELETE FROM module_submissions WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, idSubmission)
	if err != nil {
		return err
	}

	return nil
}

func (repository *moduleSubmissionsRepository) Next(ctx context.Context, tx *sql.Tx, idCourse int, idSubmission int) (entity.NextPreviousModuleSubmissions, error) {
	query := `SELECT 
				ma.id,
				c.code_course
			  FROM module_submissions ma
			  LEFT JOIN courses c ON c.id = ma.course_id 
			  WHERE ma.id > ? AND ma.course_id = ?
			  LIMIT 1`
	queryContext, err := tx.QueryContext(ctx, query, idSubmission, idCourse)
	if err != nil {
		return entity.NextPreviousModuleSubmissions{}, err
	}
	defer func(queryContext *sql.Rows) {
		err := queryContext.Close()
		if err != nil {
			return
		}
	}(queryContext)

	var ModSub entity.NextPreviousModuleSubmissions
	if queryContext.Next() {
		err := queryContext.Scan(
			&ModSub.Id,
			&ModSub.CodeCourse,
		)
		if err != nil {
			return entity.NextPreviousModuleSubmissions{}, err
		}

		return ModSub, nil
	}

	return ModSub, errors.New("submission not found")
}

func (repository *moduleSubmissionsRepository) Previous(ctx context.Context, tx *sql.Tx, idCourse int, idSubmission int) (entity.NextPreviousModuleSubmissions, error) {
	query := `SELECT 
				ma.id,
				c.code_course
			  FROM module_submissions ma
			  LEFT JOIN courses c ON c.id = ma.course_id 
			  WHERE ma.id < ? AND ma.course_id = ?
			  ORDER BY ma.id DESC
			  LIMIT 1`
	queryContext, err := tx.QueryContext(ctx, query, idSubmission, idCourse)
	if err != nil {
		return entity.NextPreviousModuleSubmissions{}, err
	}
	defer func(queryContext *sql.Rows) {
		err := queryContext.Close()
		if err != nil {
			return
		}
	}(queryContext)

	var ModSub entity.NextPreviousModuleSubmissions
	if queryContext.Next() {
		err := queryContext.Scan(
			&ModSub.Id,
			&ModSub.CodeCourse,
		)
		if err != nil {
			return entity.NextPreviousModuleSubmissions{}, err
		}

		return ModSub, nil
	}

	return ModSub, errors.New("submission not found")
}
