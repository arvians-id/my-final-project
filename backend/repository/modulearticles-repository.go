package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/rg-km/final-project-engineering-12/backend/entity"
)

type ModuleArticlesRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx, idCourse int) ([]entity.ModuleArticles, error)
	FindByModId(ctx context.Context, tx *sql.Tx, idCourse int, idArticle int) (entity.ModuleArticles, error)
	Create(ctx context.Context, tx *sql.Tx, ModArs entity.ModuleArticles) (entity.ModuleArticles, error)
	Update(ctx context.Context, tx *sql.Tx, ModArs entity.ModuleArticles, idArticle int) (entity.ModuleArticles, error)
	Delete(ctx context.Context, tx *sql.Tx, idArticle int) error
	Next(ctx context.Context, tx *sql.Tx, idCourse int, idArticle int) (entity.NextPreviousModuleArticles, error)
	Previous(ctx context.Context, tx *sql.Tx, idCourse int, idArticle int) (entity.NextPreviousModuleArticles, error)
}

type moduleArticlesRepository struct {
}

func NewModuleArticlesRepository() ModuleArticlesRepository {
	return &moduleArticlesRepository{}
}

func (repository *moduleArticlesRepository) FindAll(ctx context.Context, tx *sql.Tx, idCourse int) ([]entity.ModuleArticles, error) {
	query := `SELECT * FROM module_articles WHERE course_id = ?`
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

	var ModArs []entity.ModuleArticles
	for queryContext.Next() {
		var ModAr entity.ModuleArticles
		err := queryContext.Scan(
			&ModAr.Id,
			&ModAr.CourseId,
			&ModAr.Name,
			&ModAr.Content,
			&ModAr.Estimate,
		)
		if err != nil {
			return nil, err
		}

		ModArs = append(ModArs, ModAr)
	}

	return ModArs, nil
}

func (repository *moduleArticlesRepository) FindByModId(ctx context.Context, tx *sql.Tx, idCourse int, idArticle int) (entity.ModuleArticles, error) {
	query := `SELECT * FROM module_articles WHERE course_id = ? AND id = ?`
	queryContext, err := tx.QueryContext(ctx, query, idCourse, idArticle)
	if err != nil {
		return entity.ModuleArticles{}, err
	}
	defer func(queryContext *sql.Rows) {
		err := queryContext.Close()
		if err != nil {
			return
		}
	}(queryContext)

	var ModAr entity.ModuleArticles
	if queryContext.Next() {
		err := queryContext.Scan(
			&ModAr.Id,
			&ModAr.CourseId,
			&ModAr.Name,
			&ModAr.Content,
			&ModAr.Estimate,
		)
		if err != nil {
			return entity.ModuleArticles{}, err
		}

		return ModAr, nil
	}

	return ModAr, errors.New("article not found")
}

func (repository *moduleArticlesRepository) Create(ctx context.Context, tx *sql.Tx, ModArs entity.ModuleArticles) (entity.ModuleArticles, error) {
	query := `INSERT INTO module_articles(course_id,name,content,estimate) VALUES(?,?,?,?)`
	queryContext, err := tx.ExecContext(
		ctx,
		query,
		ModArs.CourseId,
		ModArs.Name,
		ModArs.Content,
		ModArs.Estimate,
	)
	if err != nil {
		return entity.ModuleArticles{}, err
	}

	id, err := queryContext.LastInsertId()
	if err != nil {
		return entity.ModuleArticles{}, err
	}
	ModArs.Id = int(id)

	return ModArs, nil
}

func (repository *moduleArticlesRepository) Update(ctx context.Context, tx *sql.Tx, ModArs entity.ModuleArticles, idArticle int) (entity.ModuleArticles, error) {
	query := `UPDATE module_articles SET name = ?, content = ?, estimate = ? WHERE id = ?`
	_, err := tx.ExecContext(
		ctx,
		query,
		ModArs.Name,
		ModArs.Content,
		ModArs.Estimate,
		idArticle,
	)
	if err != nil {
		return entity.ModuleArticles{}, err
	}

	return ModArs, nil
}

func (repository *moduleArticlesRepository) Delete(ctx context.Context, tx *sql.Tx, idArticle int) error {
	query := "DELETE FROM module_articles WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, idArticle)
	if err != nil {
		return err
	}

	return nil
}

func (repository *moduleArticlesRepository) Next(ctx context.Context, tx *sql.Tx, idCourse int, idArticle int) (entity.NextPreviousModuleArticles, error) {
	query := `SELECT 
				ma.id,
				c.code_course
			  FROM module_articles ma
			  LEFT JOIN courses c ON c.id = ma.course_id 
			  WHERE ma.id > ? AND ma.course_id = ?
			  LIMIT 1`
	queryContext, err := tx.QueryContext(ctx, query, idArticle, idCourse)
	if err != nil {
		return entity.NextPreviousModuleArticles{}, err
	}
	defer func(queryContext *sql.Rows) {
		err := queryContext.Close()
		if err != nil {
			return
		}
	}(queryContext)

	var ModAr entity.NextPreviousModuleArticles
	if queryContext.Next() {
		err := queryContext.Scan(
			&ModAr.Id,
			&ModAr.CodeCourse,
		)
		if err != nil {
			return entity.NextPreviousModuleArticles{}, err
		}

		return ModAr, nil
	}

	return ModAr, errors.New("article not found")
}

func (repository *moduleArticlesRepository) Previous(ctx context.Context, tx *sql.Tx, idCourse int, idArticle int) (entity.NextPreviousModuleArticles, error) {
	query := `SELECT 
				ma.id,
				c.code_course
			  FROM module_articles ma
			  LEFT JOIN courses c ON c.id = ma.course_id 
			  WHERE ma.id < ? AND ma.course_id = ?
			  ORDER BY ma.id DESC
			  LIMIT 1`
	queryContext, err := tx.QueryContext(ctx, query, idArticle, idCourse)
	if err != nil {
		return entity.NextPreviousModuleArticles{}, err
	}
	defer func(queryContext *sql.Rows) {
		err := queryContext.Close()
		if err != nil {
			return
		}
	}(queryContext)

	var ModAr entity.NextPreviousModuleArticles
	if queryContext.Next() {
		err := queryContext.Scan(
			&ModAr.Id,
			&ModAr.CodeCourse,
		)
		if err != nil {
			return entity.NextPreviousModuleArticles{}, err
		}

		return ModAr, nil
	}

	return ModAr, errors.New("article not found")
}
