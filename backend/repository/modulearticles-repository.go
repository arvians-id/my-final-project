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
	Create(ctx context.Context, tx *sql.Tx, courses entity.ModuleArticles) (entity.ModuleArticles, error)
	Update(ctx context.Context, tx *sql.Tx, courses entity.ModuleArticles) (entity.ModuleArticles, error)
	Delete(ctx context.Context, tx *sql.Tx, code string) error
}

type moduleArticlesRepository struct {
}

func NewModuleArticlesRepository() CourseRepository {
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
	query := `SELECT * FROM courses WHERE code_course = ?`
	queryContext, err := tx.QueryContext(ctx, query, code)
	if err != nil {
		return entity.Courses{}, err
	}
	defer func(queryContext *sql.Rows) {
		err := queryContext.Close()
		if err != nil {
			return
		}
	}(queryContext)

	var course entity.Courses
	if queryContext.Next() {
		err := queryContext.Scan(
			&course.Id,
			&course.Name,
			&course.CodeCourse,
			&course.Class,
			&course.Tools,
			&course.About,
			&course.Description,
			&course.CreatedAt,
			&course.UpdatedAt,
		)
		if err != nil {
			return entity.Courses{}, err
		}

		return course, nil
	}

	return course, errors.New("course not found")
}

func (repository *courseRepository) Create(ctx context.Context, tx *sql.Tx, courses entity.Courses) (entity.Courses, error) {
	query := `INSERT INTO courses(name,code_course,class,tools,about,description,created_at,updated_at) VALUES(?,?,?,?,?,?,?,?)`
	queryContext, err := tx.ExecContext(
		ctx,
		query,
		courses.Name,
		courses.CodeCourse,
		courses.Class,
		courses.Tools,
		courses.About,
		courses.Description,
		courses.CreatedAt,
		courses.UpdatedAt,
	)
	if err != nil {
		return entity.Courses{}, err
	}

	id, err := queryContext.LastInsertId()
	if err != nil {
		return entity.Courses{}, err
	}
	courses.Id = int(id)

	return courses, nil
}

func (repository *courseRepository) Update(ctx context.Context, tx *sql.Tx, courses entity.Courses) (entity.Courses, error) {
	query := `UPDATE courses SET name = ?, class = ?, tools = ?, about = ?, description = ?, updated_at = ? WHERE code_course = ?`
	_, err := tx.ExecContext(
		ctx,
		query,
		courses.Name,
		courses.Class,
		courses.Tools,
		courses.About,
		courses.Description,
		courses.UpdatedAt,
		courses.CodeCourse,
	)
	if err != nil {
		return entity.Courses{}, err
	}

	return courses, nil
}

func (repository *courseRepository) Delete(ctx context.Context, tx *sql.Tx, code string) error {
	query := "DELETE FROM courses WHERE code_course = ?"
	_, err := tx.ExecContext(ctx, query, code)
	if err != nil {
		return err
	}

	return nil
}
