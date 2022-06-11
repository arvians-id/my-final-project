package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-12/backend/model"
	"github.com/rg-km/final-project-engineering-12/backend/service"
	"net/http"
	"strconv"
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
		authorized.GET("/courses/:id", controller.FindById)
		authorized.POST("/courses", controller.Create)
		authorized.PATCH("/courses/:id", controller.Update)
		authorized.DELETE("/courses/:id", controller.Delete)
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

func (controller *CourseController) FindById(ctx *gin.Context) {
	params := ctx.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	course, err := controller.CourseService.FindById(ctx.Request.Context(), id)
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

	params := ctx.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}
	course, err := controller.CourseService.Update(ctx, request, id)
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
	params := ctx.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: err.Error(),
			Data:   nil,
		})
		return
	}

	err = controller.CourseService.Delete(ctx.Request.Context(), id)
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
