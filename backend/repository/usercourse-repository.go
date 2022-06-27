package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rg-km/final-project-engineering-12/backend/entity"
)

type UserCourseRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) ([]entity.UserCourse, error)
	FindAllCourseByUserId(ctx context.Context, tx *sql.Tx, userId int) ([]entity.StudentCourse, error)
	FindAllUserByCourseId(ctx context.Context, tx *sql.Tx, courseId int) ([]entity.UserTeacherCourse, error)
	FindByUserCourse(ctx context.Context, tx *sql.Tx, id string, course string) (entity.UserCourse, error)
	Create(ctx context.Context, tx *sql.Tx, usercourses entity.UserCourse) (entity.UserCourse, error)
	Delete(ctx context.Context, tx *sql.Tx, code1 int, code2 int) error
	FindAllStudentSubmissions(ctx context.Context, tx *sql.Tx, userId int, limit int) ([]entity.StudentSubmissions, error)
	FindAllTeacherSubmissions(ctx context.Context, tx *sql.Tx, courseId int, moduleSubmissionId int) ([]entity.TeacherSubmissions, error)
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

func (repository *usercourseRepository) FindAllCourseByUserId(ctx context.Context, tx *sql.Tx, userId int) ([]entity.StudentCourse, error) {
	query := `SELECT c.id,c.name,c.code_course,c.class FROM user_course uc
			  LEFT JOIN courses c on c.id = uc.course_id
			  WHERE uc.user_id = ?`
	queryContext, err := tx.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer func(queryContext *sql.Rows) {
		err := queryContext.Close()
		if err != nil {
			return
		}
	}(queryContext)

	var usercourses []entity.StudentCourse
	for queryContext.Next() {
		var usercourse entity.StudentCourse
		err := queryContext.Scan(
			&usercourse.IdCourse,
			&usercourse.CourseName,
			&usercourse.CourseCode,
			&usercourse.CourseClass,
		)
		if err != nil {
			return nil, err
		}

		usercourses = append(usercourses, usercourse)
	}

	return usercourses, nil
}

func (repository *usercourseRepository) FindAllUserByCourseId(ctx context.Context, tx *sql.Tx, courseId int) ([]entity.UserTeacherCourse, error) {
	query := `SELECT u.id,u.name,u.username,u.email FROM user_course uc
			  LEFT JOIN users u on u.id = uc.user_id
			  WHERE uc.course_id = ?`
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

	var usercourses []entity.UserTeacherCourse
	for queryContext.Next() {
		var usercourse entity.UserTeacherCourse
		err := queryContext.Scan(
			&usercourse.IdUser,
			&usercourse.UserName,
			&usercourse.UserUsername,
			&usercourse.UserEmail,
		)
		if err != nil {
			return nil, err
		}

		usercourses = append(usercourses, usercourse)
	}

	return usercourses, nil
}

func (repository *usercourseRepository) FindByUserCourse(ctx context.Context, tx *sql.Tx, id string, course string) (entity.UserCourse, error) {
	query := `SELECT * FROM user_course WHERE user_id = ? AND course_id = ?`
	queryContext, err := tx.QueryContext(ctx, query, id, course)
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

func (repository *usercourseRepository) FindAllStudentSubmissions(ctx context.Context, tx *sql.Tx, userId int, limit int) ([]entity.StudentSubmissions, error) {
	query := `SELECT ms.id,c.name,c.code_course,ms.name,us.grade,us.file FROM user_course uc
			  LEFT JOIN courses c on c.id = uc.course_id
			  LEFT JOIN module_submissions ms on c.id = ms.course_id
			  LEFT JOIN user_submissions us on ms.id = us.module_submission_id
			  WHERE uc.user_id = ?
			  ORDER BY us.file ASC
			  LIMIT ?`
	queryContext, err := tx.QueryContext(ctx, query, userId, limit)
	if err != nil {
		return nil, err
	}
	defer func(queryContext *sql.Rows) {
		err := queryContext.Close()
		if err != nil {
			return
		}
	}(queryContext)

	var studentSubmissions []entity.StudentSubmissions
	for queryContext.Next() {
		var studentSubmission entity.StudentSubmissions
		err := queryContext.Scan(
			&studentSubmission.IdModuleSubmission,
			&studentSubmission.CourseName,
			&studentSubmission.CodeCourse,
			&studentSubmission.ModuleSubmissionName,
			&studentSubmission.Grade,
			&studentSubmission.File,
		)
		if err != nil {
			return nil, err
		}

		studentSubmissions = append(studentSubmissions, studentSubmission)
	}

	return studentSubmissions, nil
}

func (repository *usercourseRepository) FindAllTeacherSubmissions(ctx context.Context, tx *sql.Tx, courseId int, moduleSubmissionId int) ([]entity.TeacherSubmissions, error) {
	query := `SELECT us.id,u.name,ms.name,us.grade,us.file FROM user_course uc
			  LEFT JOIN users u on u.id = uc.user_id
			  LEFT JOIN courses c on c.id = uc.course_id
			  LEFT JOIN module_submissions ms on c.id = ms.course_id
			  LEFT JOIN user_submissions us on u.id = us.user_id
			  WHERE c.id = ? AND ms.id = ?`
	queryContext, err := tx.QueryContext(ctx, query, courseId, moduleSubmissionId)
	if err != nil {
		return nil, err
	}
	defer func(queryContext *sql.Rows) {
		err := queryContext.Close()
		if err != nil {
			return
		}
	}(queryContext)

	var teacherSubmissions []entity.TeacherSubmissions
	for queryContext.Next() {
		var studentSubmission entity.TeacherSubmissions
		err := queryContext.Scan(
			&studentSubmission.IdUserSubmission,
			&studentSubmission.UserName,
			&studentSubmission.ModuleSubmissionName,
			&studentSubmission.Grade,
			&studentSubmission.File,
		)
		if err != nil {
			return nil, err
		}

		teacherSubmissions = append(teacherSubmissions, studentSubmission)
	}

	return teacherSubmissions, nil
}
