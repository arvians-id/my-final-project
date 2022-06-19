package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/rg-km/final-project-engineering-12/backend/entity"
)

type UserCourseRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) ([]entity.UserCourse, error)
	FindByUserId(ctx context.Context, tx *sql.Tx, id string) (entity.UserCourse, error)
	Create(ctx context.Context, tx *sql.Tx, usercourses entity.UserCourse, user entity.Users, course entity.Courses) (entity.UserCourse, error)
	Delete(ctx context.Context, tx *sql.Tx, code string) error
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

func (repository *usercourseRepository) Create(ctx context.Context, tx *sql.Tx, usercourses entity.UserCourse, user entity.Users, course entity.Courses) (entity.UserCourse, error) {
	var iduser, idcourse int

	database, err := sql.Open("sqlite3", "./teenager.db")

	if err != nil {
		return entity.UserCourse{}, err
	}

	defer database.Close()
	rows1 := database.QueryRow("SELECT id FROM users WHERE username = ?", user.Username)
	rows1.Scan(&iduser)
	rows2 := database.QueryRow("SELECT id FROM courses WHERE code_course = ?", course.CodeCourse)
	rows2.Scan(&idcourse)

	query := `INSERT INTO user_course(user_id, course_id) VALUES(?,?)`
	queryContext, err := tx.ExecContext(
		ctx,
		query,
		iduser,
		idcourse,
	)
	if err != nil {
		return entity.UserCourse{}, err
	}

	f, err := queryContext.LastInsertId()
	if err != nil {
		return entity.UserCourse{}, err
	}
	fmt.Println(f)

	return usercourses, nil
}

func (repository *usercourseRepository) Delete(ctx context.Context, tx *sql.Tx, code string) error {
	query := "DELETE FROM user_course WHERE user_id = ?"
	_, err := tx.ExecContext(ctx, query, code)
	if err != nil {
		return err
	}

	return nil
}
