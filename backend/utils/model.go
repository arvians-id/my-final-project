package utils

import (
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
		IsActive:    course.IsActive,
	}
}

func ToModuleArticlesResponse(ModArs entity.ModuleArticles) model.GetModuleArticlesResponse {
	return model.GetModuleArticlesResponse{
		Id:       ModArs.Id,
		CourseId: ModArs.CourseId,
		Name:     ModArs.Name,
		Content:  ModArs.Content,
		Estimate: ModArs.Estimate,
	}
}

func ToModuleArticlesNextPreviousResponse(ModArs entity.NextPreviousModuleArticles) model.GetNextPreviousArticlesResponse {
	return model.GetNextPreviousArticlesResponse{
		Id:         ModArs.Id,
		CodeCourse: ModArs.CodeCourse,
	}
}

func ToModuleSubmissionsNextPreviousResponse(ModSubs entity.NextPreviousModuleSubmissions) model.GetNextPreviousSubmissionsResponse {
	return model.GetNextPreviousSubmissionsResponse{
		Id:         ModSubs.Id,
		CodeCourse: ModSubs.CodeCourse,
	}
}

func ToModuleSubmissionsResponse(modsub entity.ModuleSubmissions) model.GetModuleSubmissionsResponse {
	return model.GetModuleSubmissionsResponse{
		Id:          modsub.Id,
		CourseId:    modsub.CourseId,
		Name:        modsub.Name,
		Description: modsub.Description,
		Deadline:    modsub.Deadline,
	}
}

func ToUserSubmissionsResponse(userSubmission entity.UserSubmissions) model.GetUserSubmissionsResponse {
	return model.GetUserSubmissionsResponse{
		Id:                 userSubmission.Id,
		UserId:             userSubmission.UserId,
		ModuleSubmissionId: userSubmission.ModuleSubmissionId,
		File:               userSubmission.File,
		Grade:              userSubmission.Grade,
	}
}

func ToUserCourseResponse(usercourse entity.UserCourse) model.GetUserCourseResponse {
	return model.GetUserCourseResponse{
		UserId:   usercourse.UserId,
		CourseId: usercourse.CourseId,
	}
}

func ToQuestionResponse(question entity.Questions) model.GetQuestionResponse {
	return model.GetQuestionResponse{
		Id:          question.Id,
		ModuleId:    question.ModuleId,
		UserId:      question.UserId,
		Title:       question.Title,
		Tags:        question.Tags,
		Description: question.Description,
		CreatedAt:   question.CreatedAt,
		UpdatedAt:   question.UpdatedAt,
	}
}

func ToAnswerResponse(answer entity.Answers) model.GetAnswerResponse {
	return model.GetAnswerResponse{
		Id:          answer.Id,
		QuestionId:  answer.QuestionId,
		UserId:      answer.UserId,
		Description: answer.Description,
		CreatedAt:   answer.CreatedAt,
		UpdatedAt:   answer.UpdatedAt,
	}
}

func ToStudentSubmissionsResponse(submission entity.StudentSubmissions) model.GetStudentSubmissionsResponse {
	return model.GetStudentSubmissionsResponse{
		IdModuleSubmission:   submission.IdModuleSubmission,
		NameCourse:           submission.CourseName,
		NameModuleSubmission: submission.ModuleSubmissionName,
		Grade:                submission.Grade,
		File:                 submission.File,
	}
}

func ToTeacherSubmissionsResponse(submission entity.TeacherSubmissions) model.GetTeacherSubmissionsResponse {
	return model.GetTeacherSubmissionsResponse{
		IdUserSubmission:     submission.IdUserSubmission,
		UserName:             submission.UserName,
		ModuleSubmissionName: submission.ModuleSubmissionName,
		Grade:                submission.Grade,
		File:                 submission.File,
	}
}
