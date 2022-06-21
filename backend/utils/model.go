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
