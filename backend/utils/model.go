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

func ToModuleSubmissionsResponse(modsub entity.ModuleSubmissions) model.GetModuleSubmissionsResponse {
	return model.GetModuleSubmissionsResponse{
		Id:       modsub.Id,
		ModuleId: modsub.ModuleId,
		File:     modsub.File,
		Type:     modsub.Type,
		MaxSize:  modsub.MaxSize,
	}
}
