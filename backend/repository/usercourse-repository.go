package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rg-km/final-project-engineering-12/backend/entity"
)

type UserCourseRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) ([]entity.UserCourse, error)
	FindByUserId(ctx context.Context, tx *sql.Tx, id string) (entity.UserCourse, error)
	Create(ctx context.Context, tx *sql.Tx, usercourses entity.UserCourse) (entity.UserCourse, error)
	Delete(ctx context.Context, tx *sql.Tx, code1 int, code2 int) error
}

type usercourseRepository struct {
}

func NewUserCourseRepository() UserCourseRepository {
	return &usercourseRepository{}
}

func (repository *usercourseRepository) FindAll(ctx context.Context, tx *sql.Tx) ([]entity.UserCourse, error) {
	query := `SELECT * FROM user_course ORDER BY user_id DESC`
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

	var usercourses []entity.UserCourse
	for queryContext.Next() {
		var usercourse entity.UserCourse
		err := queryContext.Scan(
			&usercourse.UserId,
			&usercourse.CourseId,
		)
		if err != nil {
			return nil, err
		}

		usercourses = append(usercourses, usercourse)
	}

	return usercourses, nil
}

func (repository *usercourseRepository) FindByUserId(ctx context.Context, tx *sql.Tx, id string) (entity.UserCourse, error) {
	query := `SELECT * FROM user_course WHERE user_id = ?`
	queryContext, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return entity.UserCourse{}, err
	}
	defer func(queryContext *sql.Rows) {
		err := queryContext.Close()
		if err != nil {
			return
		}
	}(queryContext)

	var usercourse entity.UserCourse
	if queryContext.Next() {
		err := queryContext.Scan(
			&usercourse.UserId,
			&usercourse.CourseId,
		)
		if err != nil {
			return entity.UserCourse{}, err
		}

		return usercourse, nil
	}

	return usercourse, errors.New("course not found")
}

func (repository *usercourseRepository) Create(ctx context.Context, tx *sql.Tx, usercourses entity.UserCourse) (entity.UserCourse, error) {
	_, err := tx.ExecContext(ctx, "INSERT INTO user_course (user_id, course_id) VALUES(?,?)", usercourses.UserId, usercourses.CourseId)
	if err != nil {
		return entity.UserCourse{}, err
	}

	return usercourses, nil
}

func (repository *usercourseRepository) Delete(ctx context.Context, tx *sql.Tx, code1 int, code2 int) error {
	query := "DELETE FROM user_course WHERE user_id = ? AND course_id = ?"
	_, err := tx.ExecContext(ctx, query, code1, code2)
	if err != nil {
		return err
	}

	return nil
}
