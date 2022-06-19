package utils

import (
	"fmt"
	"github.com/rg-km/final-project-engineering-12/backend/entity"
	"github.com/rg-km/final-project-engineering-12/backend/model"
)

func ToCourseResponse(course entity.Courses) model.GetCourseResponse {
	return model.GetCourseResponse{
		Id:          course.Id,
		Name:        course.Name,
		CodeCourse:  course.CodeCourse,
		Class:       course.Class,
		Tools:       course.Tools,
		About:       course.About,
		Description: course.Description,
		CreatedAt:   course.CreatedAt,
		UpdatedAt:   course.UpdatedAt,
	}
}

func ToUserCourseResponse(usercourse entity.UserCourse) model.GetUserCourseResponse {
	return model.GetUserCourseResponse{
		UserId:   		usercourse.UserId,
		CourseId:   	usercourse.CourseId,
	}
}

func ToQuestionResponse(question entity.Questions) model.GetQuestionResponse {
	fmt.Println(question)
	return model.GetQuestionResponse{
		Id:         	question.Id,
		ModuleId:   	question.ModuleId,
		UserId:   		question.UserId,
		Title:  			question.Title,
		Tags:       	question.Tags,
		Description: 	question.Description,
		CreatedAt:   	question.CreatedAt,
		UpdatedAt:   	question.UpdatedAt,
	}
}

func ToAnswerResponse(answer entity.Answers) model.GetAnswerResponse {
	return model.GetAnswerResponse{
		Id:         	answer.Id,
		QuestionId:   answer.QuestionId,
		UserId:   		answer.UserId,
		Description: 	answer.Description,
		CreatedAt:   	answer.CreatedAt,
		UpdatedAt:   	answer.UpdatedAt,
	}
}