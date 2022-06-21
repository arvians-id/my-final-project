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
