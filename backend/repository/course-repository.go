package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/rg-km/final-project-engineering-12/backend/entity"
)

type CourseRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx, status bool, limit int) ([]entity.Courses, error)
	FindById(ctx context.Context, tx *sql.Tx, id int) (entity.Courses, error)
	Create(ctx context.Context, tx *sql.Tx, courses entity.Courses) (entity.Courses, error)
	Update(ctx context.Context, tx *sql.Tx, courses entity.Courses, id int) (entity.Courses, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	ChangeActiveCourse(ctx context.Context, tx *sql.Tx, status bool, id int) error
}

type courseRepository struct {
}

func NewCourseRepository() CourseRepository {
	return &courseRepository{}
}

func (repository *courseRepository) FindAll(ctx context.Context, tx *sql.Tx, status bool, limit int) ([]entity.Courses, error) {
	query := `SELECT * FROM courses WHERE is_active = ? ORDER BY created_at DESC LIMIT ?`
	queryContext, err := tx.QueryContext(ctx, query, status, limit)
	if err != nil {
		return nil, err
	}
	defer func(queryContext *sql.Rows) {
		err := queryContext.Close()
		if err != nil {
			return
		}
	}(queryContext)

	var courses []entity.Courses
	for queryContext.Next() {
		var course entity.Courses
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
			&course.IsActive,
		)
		if err != nil {
			return nil, err
		}

		courses = append(courses, course)
	}

	return courses, nil
}

func (repository *courseRepository) FindById(ctx context.Context, tx *sql.Tx, id int) (entity.Courses, error) {
	query := `SELECT * FROM courses WHERE id = ?`
	queryContext, err := tx.QueryContext(ctx, query, id)
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
			&course.IsActive,
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
		courses.IsActive,
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

func (repository *courseRepository) Update(ctx context.Context, tx *sql.Tx, courses entity.Courses, id int) (entity.Courses, error) {
	query := `UPDATE courses SET name = ?, class = ?, tools = ?, about = ?, description = ?, updated_at = ? WHERE id = ?`
	_, err := tx.ExecContext(
		ctx,
		query,
		courses.Name,
		courses.Class,
		courses.Tools,
		courses.About,
		courses.Description,
		courses.UpdatedAt,
		id,
	)
	if err != nil {
		return entity.Courses{}, err
	}

	return courses, nil
}

func (repository *courseRepository) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM courses WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (repository *courseRepository) ChangeActiveCourse(ctx context.Context, tx *sql.Tx, status bool, id int) error {
	query := `UPDATE courses SET is_active = ? WHERE id = ?`
	_, err := tx.ExecContext(
		ctx,
		query,
		status,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}
