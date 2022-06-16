package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/rg-km/final-project-engineering-12/backend/entity"
)

type ModuleRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx, limit int) ([]entity.Modules, error)
	FindAllByRelation(ctx context.Context, tx *sql.Tx, courseId int) ([]entity.Modules, error)
	FindById(ctx context.Context, tx *sql.Tx, courseId int, id int) (entity.Modules, error)
	Create(ctx context.Context, tx *sql.Tx, modules entity.Modules) (entity.Modules, error)
	Update(ctx context.Context, tx *sql.Tx, modules entity.Modules, id int) (entity.Modules, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
}

type moduleRepository struct {
}

func NewModuleRepository() ModuleRepository {
	return &moduleRepository{}
}

func (repository *moduleRepository) FindAll(ctx context.Context, tx *sql.Tx, limit int) ([]entity.Modules, error) {
	query := `SELECT * FROM modules LIMIT ?`
	queryContext, err := tx.QueryContext(ctx, query, limit)
	if err != nil {
		return nil, err
	}
	defer func(queryContext *sql.Rows) {
		err := queryContext.Close()
		if err != nil {
			return
		}
	}(queryContext)

	var modules []entity.Modules
	for queryContext.Next() {
		var module entity.Modules
		err := queryContext.Scan(
			&module.Id,
			&module.CourseId,
			&module.Name,
			&module.IsLocked,
			&module.Estimate,
			&module.Deadline,
			&module.Grade,
		)
		if err != nil {
			return nil, err
		}

		modules = append(modules, module)
	}

	return modules, nil
}

func (repository *moduleRepository) FindAllByRelation(ctx context.Context, tx *sql.Tx, courseId int) ([]entity.Modules, error) {
	query := `
				SELECT 
					m.*
				FROM modules m
				LEFT JOIN courses c ON c.id = m.course_id
				WHERE c.id = ?`
	queryContext, err := tx.QueryContext(ctx, query, courseId)
	if err != nil {
		return nil, err
	}
	defer func(queryContext *sql.Rows) {
		err := queryContext.Close()
		if err != nil {
			return
		}
	}(queryContext)

	var modules []entity.Modules
	for queryContext.Next() {
		var module entity.Modules
		err := queryContext.Scan(
			&module.Id,
			&module.CourseId,
			&module.Name,
			&module.IsLocked,
			&module.Estimate,
			&module.Deadline,
			&module.Grade,
		)
		if err != nil {
			return nil, err
		}

		modules = append(modules, module)
	}

	return modules, nil
}

func (repository *moduleRepository) FindById(ctx context.Context, tx *sql.Tx, courseId int, id int) (entity.Modules, error) {
	query := `
				SELECT 
					m.*
				FROM modules m
				LEFT JOIN courses c ON c.id = m.course_id
				WHERE c.id = ?
				AND m.id = ?`
	queryContext, err := tx.QueryContext(ctx, query, courseId, id)
	if err != nil {
		return entity.Modules{}, err
	}
	defer func(queryContext *sql.Rows) {
		err := queryContext.Close()
		if err != nil {
			return
		}
	}(queryContext)

	var module entity.Modules
	if queryContext.Next() {
		err := queryContext.Scan(
			&module.Id,
			&module.CourseId,
			&module.Name,
			&module.IsLocked,
			&module.Estimate,
			&module.Deadline,
			&module.Grade,
		)
		if err != nil {
			return entity.Modules{}, err
		}

		return module, nil
	}

	return module, errors.New("module not found")
}

func (repository *moduleRepository) Create(ctx context.Context, tx *sql.Tx, modules entity.Modules) (entity.Modules, error) {
	query := `INSERT INTO modules(course_id,name,is_locked,estimate,deadline) VALUES(?,?,?,?,?)`
	queryContext, err := tx.ExecContext(
		ctx,
		query,
		modules.CourseId,
		modules.Name,
		modules.IsLocked,
		modules.Estimate,
		modules.Deadline,
	)
	if err != nil {
		return entity.Modules{}, err
	}

	id, err := queryContext.LastInsertId()
	if err != nil {
		return entity.Modules{}, err
	}
	modules.Id = int(id)

	return modules, nil
}

func (repository *moduleRepository) Update(ctx context.Context, tx *sql.Tx, modules entity.Modules, id int) (entity.Modules, error) {
	query := `UPDATE modules SET course_id = ?, name = ?, is_locked = ?, estimate = ?, deadline = ? WHERE id = ?`
	_, err := tx.ExecContext(
		ctx,
		query,
		modules.CourseId,
		modules.Name,
		modules.IsLocked,
		modules.Estimate,
		modules.Deadline,
		id,
	)
	if err != nil {
		return entity.Modules{}, err
	}

	return modules, nil
}

func (repository *moduleRepository) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM modules WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
