package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rg-km/final-project-engineering-12/backend/entity"
)

type ModuleArticlesRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) ([]entity.ModuleArticles, error)
	FindByModId(ctx context.Context, tx *sql.Tx, code string) (entity.ModuleArticles, error)
	Create(ctx context.Context, tx *sql.Tx, ModArs entity.ModuleArticles) (entity.ModuleArticles, error)
	Update(ctx context.Context, tx *sql.Tx, ModArs entity.ModuleArticles) (entity.ModuleArticles, error)
	Delete(ctx context.Context, tx *sql.Tx, code string) error
}

type moduleArticlesRepository struct {
}

func NewModuleArticlesRepository() ModuleArticlesRepository {
	return &moduleArticlesRepository{}
}

func (repository *moduleArticlesRepository) FindAll(ctx context.Context, tx *sql.Tx) ([]entity.ModuleArticles, error) {
	query := `SELECT * FROM module_articles ORDER BY module_id ASC`
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

	var ModArs []entity.ModuleArticles
	for queryContext.Next() {
		var ModAr entity.ModuleArticles
		err := queryContext.Scan(
			&ModAr.Id,
			&ModAr.ModuleId,
			&ModAr.Content,
		)
		if err != nil {
			return nil, err
		}

		ModArs = append(ModArs, ModAr)
	}

	return ModArs, nil
}

func (repository *moduleArticlesRepository) FindByModId(ctx context.Context, tx *sql.Tx, code string) (entity.ModuleArticles, error) {
	query := `SELECT * FROM module_articles WHERE module_id = ?`
	queryContext, err := tx.QueryContext(ctx, query, code)
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
			&ModAr.ModuleId,
			&ModAr.Content,
		)
		if err != nil {
			return entity.ModuleArticles{}, err
		}

		return ModAr, nil
	}

	return ModAr, errors.New("article not found")
}

func (repository *moduleArticlesRepository) Create(ctx context.Context, tx *sql.Tx, ModArs entity.ModuleArticles) (entity.ModuleArticles, error) {
	query := `INSERT INTO module_articles(module_id,content) VALUES(?,?)`
	queryContext, err := tx.ExecContext(
		ctx,
		query,
		ModArs.ModuleId,
		ModArs.Content,
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

func (repository *moduleArticlesRepository) Update(ctx context.Context, tx *sql.Tx, ModArs entity.ModuleArticles) (entity.ModuleArticles, error) {
	query := `UPDATE module_articles SET content = ? WHERE module_id = ?`
	_, err := tx.ExecContext(
		ctx,
		query,
		ModArs.Content,
		ModArs.ModuleId,
	)
	if err != nil {
		return entity.ModuleArticles{}, err
	}

	return ModArs, nil
}

func (repository *moduleArticlesRepository) Delete(ctx context.Context, tx *sql.Tx, code string) error {
	query := "DELETE FROM module_articles WHERE module_id = ?"
	_, err := tx.ExecContext(ctx, query, code)
	if err != nil {
		return err
	}

	return nil
}
