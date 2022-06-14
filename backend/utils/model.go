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
	}
}

func ToModuleArticlesResponse(ModArs entity.ModuleArticles) model.GetModuleArticlesResponse {
	return model.GetModuleArticlesResponse{
		Id:       ModArs.Id,
		ModuleId: ModArs.ModuleId,
		Content:  ModArs.Content,
	}
}
