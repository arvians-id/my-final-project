package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/rg-km/final-project-engineering-12/backend/entity"
)

type ModuleSubmissionsRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) ([]entity.ModuleSubmissions, error)
	FindByModId(ctx context.Context, tx *sql.Tx, code string) (entity.ModuleSubmissions, error)
	Create(ctx context.Context, tx *sql.Tx, modsub entity.ModuleSubmissions) (entity.ModuleSubmissions, error)
	Update(ctx context.Context, tx *sql.Tx, modsub entity.ModuleSubmissions) (entity.ModuleSubmissions, error)
	Delete(ctx context.Context, tx *sql.Tx, code string) error
}

type moduleSubmissionsRepository struct {
}

func NewModuleSubmissionsRepository() ModuleSubmissionsRepository {
	return &moduleSubmissionsRepository{}
}

func (repository *moduleSubmissionsRepository) FindAll(ctx context.Context, tx *sql.Tx) ([]entity.ModuleSubmissions, error) {
	query := `SELECT * FROM module_submissions ORDER BY module_id ASC`
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

	var modsubs []entity.ModuleSubmissions
	for queryContext.Next() {
		var modsub entity.ModuleSubmissions
		err := queryContext.Scan(
			&modsub.Id,
			&modsub.ModuleId,
			&modsub.File,
			&modsub.Type,
			&modsub.MaxSize,
		)
		if err != nil {
			return nil, err
		}

		modsubs = append(modsubs, modsub)
	}

	return modsubs, nil
}

func (repository *moduleSubmissionsRepository) FindByModId(ctx context.Context, tx *sql.Tx, code string) (entity.ModuleSubmissions, error) {
	query := `SELECT * FROM module_submissions WHERE module_id = ?`
	queryContext, err := tx.QueryContext(ctx, query, code)
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
			&modsub.ModuleId,
			&modsub.File,
			&modsub.Type,
			&modsub.MaxSize,
		)
		if err != nil {
			return entity.ModuleSubmissions{}, err
		}

		return modsub, nil
	}

	return modsub, errors.New("submissions not found")
}

func (repository *moduleSubmissionsRepository) Create(ctx context.Context, tx *sql.Tx, modsub entity.ModuleSubmissions) (entity.ModuleSubmissions, error) {
	query := `INSERT INTO module_submissions(module_id, file, type, max_size) VALUES(?,?,?,?)`
	queryContext, err := tx.ExecContext(
		ctx,
		query,
		&modsub.ModuleId,
		&modsub.File,
		&modsub.Type,
		&modsub.MaxSize,
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

func (repository *moduleSubmissionsRepository) Update(ctx context.Context, tx *sql.Tx, modsub entity.ModuleSubmissions) (entity.ModuleSubmissions, error) {
	query := `UPDATE module_submissions SET file = ?, type = ?, max_size = ? WHERE module_id = ?`
	_, err := tx.ExecContext(
		ctx,
		query,
		&modsub.File,
		&modsub.Type,
		&modsub.MaxSize,
		&modsub.ModuleId,
	)
	if err != nil {
		return entity.ModuleSubmissions{}, err
	}

	return modsub, nil
}

func (repository *moduleSubmissionsRepository) Delete(ctx context.Context, tx *sql.Tx, code string) error {
	query := "DELETE FROM module_submissions WHERE module_id = ?"
	_, err := tx.ExecContext(ctx, query, code)
	if err != nil {
		return err
	}

	return nil
}
