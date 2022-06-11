package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/service"
	"net/http"
)

type CourseController struct {
	CourseService service.CourseService
}

func NewCourseController(courseService *service.CourseService) *CourseController {
	return &CourseController{
		CourseService: *courseService,
	}
}

func (controller *CourseController) Route(router *gin.Engine) *gin.Engine {
	authorized := router.Group("/api")
	{
		authorized.GET("/courses", controller.FindAll)
		authorized.GET("/courses/:code", controller.FindByCode)
		authorized.POST("/courses", controller.Create)
		authorized.PATCH("/courses/:code", controller.Update)
		authorized.DELETE("/courses/:code", controller.Delete)
	}

	return router
}

func (controller *CourseController) FindAll(ctx *gin.Context) {
	courses, err := controller.CourseService.FindAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   courses,
	})
}

func (controller *CourseController) FindByCode(ctx *gin.Context) {
	code := ctx.Param("code")
	course, err := controller.CourseService.FindByCourse(ctx.Request.Context(), code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   course,
	})
}

func (controller *CourseController) Create(ctx *gin.Context) {
	var request model.CreateCourseRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	course, err := controller.CourseService.Create(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Status: "course successfully created",
		Data:   course,
	})
}

func (controller *CourseController) Update(ctx *gin.Context) {
	var request model.UpdateCourseRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	code := ctx.Param("code")
	course, err := controller.CourseService.Update(ctx, request, code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Status: "course successfully updated",
		Data:   course,
	})
}

func (controller *CourseController) Delete(ctx *gin.Context) {
	code := ctx.Param("code")

	err := controller.CourseService.Delete(ctx.Request.Context(), code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Status: "course successfully deleted",
		Data:   nil,
	})
}
